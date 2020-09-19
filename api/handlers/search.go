package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/dbops"
	"github.com/piggona/fundingsView/api/defs"
	"github.com/piggona/fundingsView/api/utils/log"
)

const (
	elasticsearchError = "001"
	requestEncodeError = "000"
)

func HomePage(c *gin.Context) {
	var err error
	// 主页元素统计-总投资数
	statisticAmount, err := dbops.GetBasicStatisticsAmount()
	statisticAvg, err := dbops.GetBasicStatisticsAvg()
	// 热点-产业 & 主页元素-排名-按金额排序（产业排序）
	hotIndu, err := dbops.GetBasicRankAmount(0, 20)
	// 主分类资金分布饼图
	basicPie, err := dbops.GetBasicPie()
	// 主页元素-解析-技术
	techAnalysis, err := dbops.GetBasicAnalysisTech()
	if err != nil {
		sendErrorResponse(c, defs.ErrorResponse{
			HTTPSc: http.StatusInternalServerError,
			Error: defs.Err{
				Error:     fmt.Sprintf("es search error: %s", err),
				ErrorCode: elasticsearchError,
			},
		})
	}
	// 合成返回
	resPie := pieParser(basicPie)
	resCloud := cloudParser(hotIndu)
	fundRank := rankAmountParser(hotIndu)
	treeNode := treeTechParser(techAnalysis)

	result := &defs.BasicResponse{
		Statistic: defs.Statistic{
			FundingsAmount: *statisticAmount,
			AvgInvested:    *statisticAvg,
		},
		ResPie:    *resPie,
		ResCloud:  *resCloud,
		FundRank:  *fundRank,
		TreeNodes: *treeNode,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   result,
	})

}

func pieParser(basicPie []*dbops.BasicPie) *defs.ResPieResponse {
	legends := make([]string, len(basicPie))
	series := make([]*defs.SeriesData, len(basicPie))
	for id, pie := range basicPie {
		series[id] = &defs.SeriesData{
			Value: pie.Value,
			Name:  pie.Name,
			Key:   pie.Key,
		}
		legends[id] = pie.Key
	}
	return &defs.ResPieResponse{
		Legend: legends,
		Series: series,
	}
}

func cloudParser(hotIndu []*dbops.BasicRankAmountResult) *defs.ResCloudResponse {
	data := make([]*defs.SeriesData, len(hotIndu))
	for id, indu := range hotIndu {
		data[id].Key = indu.Key
		data[id].Value = indu.Value
	}
	return &defs.ResCloudResponse{
		Data: data,
	}
}

func rankAmountParser(hotIndu []*dbops.BasicRankAmountResult) *defs.FundRankResponse {
	data := make([]*defs.FundRankData, len(hotIndu))
	orderType := "amount"
	for id, res := range hotIndu {
		data[id].Name = res.Key
		data[id].Money = strconv.Itoa(res.Value)
		data[id].Rank = strconv.Itoa(id)
	}
	return &defs.FundRankResponse{
		OrderType: orderType,
		Data:      data,
	}
}

type GrowthRank struct {
	Key  string
	Rate float64
}

type GrowthRankList []*GrowthRank

func (g GrowthRankList) Len() int {
	return len(g)
}

func (g GrowthRankList) Less(i, j int) bool {
	return g[i].Rate < g[j].Rate
}

func (g GrowthRankList) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func rankGrowthParser(growthIndu []*dbops.BasicRankGrowthResult) *defs.FundRankResponse {
	data := make([]*defs.FundRankData, len(growthIndu))
	orderGrowth := make([]*GrowthRank, len(growthIndu))
	orderType := "growth"
	for id, res := range growthIndu {
		orderGrowth[id] = &GrowthRank{
			Key:  res.Key,
			Rate: getRate(res.DateValue),
		}
	}
	sort.Sort(GrowthRankList(orderGrowth))
	for id, growth := range orderGrowth {
		data[id].Rank = strconv.Itoa(id)
		data[id].Name = growth.Key
		data[id].Money = strconv.FormatFloat(growth.Rate-1.0, 'f', 2, 64)
	}
	return &defs.FundRankResponse{
		OrderType: orderType,
		Data:      data,
	}

}

func getRate(dateVal map[string]int) float64 {
	var maxTime *time.Time
	var minTime *time.Time
	var maxVal int
	var minVal int
	for t, val := range dateVal {
		curTime, err := time.Parse("01/02/2006", t)
		if err != nil {
			log.Error("time parse failed: %s.string is: %s", err, t)
			return 0
		}
		if maxTime == nil || curTime.After(*maxTime) {
			maxTime = &curTime
			maxVal = val
		}
		if minTime == nil || curTime.Before(*minTime) {
			minTime = &curTime
			minVal = val
		}
	}
	interval := maxTime.Year() - minTime.Year()
	return float64(maxVal-minVal) / float64(interval)
}

func treeTechParser(techAnalysis []*dbops.BasicAnalysisTechResult) *defs.TreeNodeResponse {
	data := make([]*defs.TreeData, len(techAnalysis))
	for id, div := range techAnalysis {
		data[id].Key = div.Key
		data[id].Title = div.Name
		children := make([]*defs.TreeData, len(div.Techs))
		for i, tech := range div.Techs {
			children[i].Title = tech.Tech
			children[i].Key = tech.Tech
		}
		data[id].Children = children
	}
	return &defs.TreeNodeResponse{
		Data: data,
	}
}

func treeInduParser(induAnalysis []*dbops.BasicAnalysisInduResult) *defs.TreeNodeResponse {
	data := make([]*defs.TreeData, len(induAnalysis))
	for id, div := range induAnalysis {
		data[id].Key = div.Key
		data[id].Title = div.Name
		children := make([]*defs.TreeData, len(div.Indus))
		for i, tech := range div.Indus {
			children[i].Title = tech.Indu
			children[i].Key = tech.Indu
		}
		data[id].Children = children
	}
	return &defs.TreeNodeResponse{
		Data: data,
	}
}

func SwitchRank(c *gin.Context) {
	activeAmount := c.Param("activeAmount")
	switch activeAmount {
	case "amount":
		rankAmount, err := dbops.GetBasicRankAmount(0, 20)
		if err != nil {
			log.Error("dbops get basic rank amount error: %s", err)
			sendErrorResponse(c, defs.ErrorDBError)
			return
		}
		resp := rankAmountParser(rankAmount)
		sendNormalResponse(c, defs.NormalResp{
			HttpSc: http.StatusOK,
			Resp:   resp,
		})

	case "growth":
		// 主页元素-排名-按增长率排序
		rankGrowth, err := dbops.GetBasicRankGrowth(0, 10)
		if err != nil {
			log.Error("dbops get basic rank growth error: %s", err)
			sendErrorResponse(c, defs.ErrorDBError)
			return
		}
		resp := rankGrowthParser(rankGrowth)
		sendNormalResponse(c, defs.NormalResp{
			HttpSc: http.StatusOK,
			Resp:   resp,
		})
	}
	return
}

func SwitchTech(c *gin.Context) {
	activeTech := c.Param("activeTech")
	switch activeTech {
	case "tech":
		techAnalysis, err := dbops.GetBasicAnalysisTech()
		if err != nil {
			log.Error("dbops get basic analysis tech error: %s", err)
			sendErrorResponse(c, defs.ErrorDBError)
			return
		}
		resp := treeTechParser(techAnalysis)
		sendNormalResponse(c, defs.NormalResp{
			HttpSc: http.StatusOK,
			Resp:   resp,
		})
	case "indu":
		induAnalysis, err := dbops.GetBasicAnalysisIndu()
		if err != nil {
			log.Error("dbops get basic analysis indu error: %s", err)
			sendErrorResponse(c, defs.ErrorDBError)
			return
		}
		resp := treeInduParser(induAnalysis)
		sendNormalResponse(c, defs.NormalResp{
			HttpSc: http.StatusOK,
			Resp:   resp,
		})
	}
	return
}

func GetSearch(c *gin.Context) {
	searchReq := &defs.SearchRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(searchReq)
	if err != nil {
		log.Error("read request body error: %s", err)
		sendErrorResponse(c, defs.ErrorResponse{
			HTTPSc: http.StatusInternalServerError,
			Error: defs.Err{
				Error:     fmt.Sprintf("%s", err),
				ErrorCode: requestEncodeError,
			},
		})
		return
	}
	dbSearch := &dbops.SearchParams{
		From: strconv.Itoa(searchReq.From),
	}
	for _, element := range searchReq.Array {
		switch element.Attr {
		case "basic":
			dbSearch.Title = element.Query
			dbSearch.Description = element.Query
			dbSearch.TitleOp = element.Logic
			dbSearch.DescriptionOp = element.Logic
		case "Investigator":
			dbSearch.Institution = element.Query
			dbSearch.InstitutionOp = element.Logic
		case "Institution":
			dbSearch.Organization = element.Query
			dbSearch.OrganizationOp = element.Logic
		case "Technology":
			dbSearch.Technology = element.Query
			dbSearch.TechnologyOp = element.Logic
		case "Industry":
			dbSearch.Industries = element.Query
			dbSearch.IndustriesOp = element.Logic
		case "Date":
			dateRange := strings.Split(element.Query, " ")
			dbSearch.DateFrom, _ = time.Parse("2006-01-02", dateRange[0])
			dbSearch.DateTo, _ = time.Parse("2006-01-02", dateRange[1])
		}
	}
	elements, err := dbops.GetMultiSearch(dbSearch)
	if err != nil {
		log.Error("multisearch error: %s", err)
		sendErrorResponse(c, defs.ErrorResponse{
			HTTPSc: http.StatusInternalServerError,
			Error: defs.Err{
				Error:     fmt.Sprintf("%s", err),
				ErrorCode: elasticsearchError,
			},
		})
		return
	}
	resp := &defs.SearchResponse{
		TotalResults: len(elements),
		Data:         make([]*defs.SearchResultBucket, len(elements)),
	}
	for id, element := range elements {
		bucket := &defs.SearchResultBucket{
			ID:     strconv.Itoa(id),
			Title:  element.AwardTitle,
			Amount: strconv.Itoa(element.AwardAmount),
			UUID:   element.AwardID,
			Start:  element.AwardEffecticeDate,
			End:    element.AwardExpirationDate,
		}
		resp.Data[id] = bucket
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}
