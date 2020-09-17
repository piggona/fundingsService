package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

// func HomePage(c *gin.Context) {
// 	var err error
// 	// 主页元素统计-总投资数
// 	statisticAmount, err := dbops.GetBasicStatisticsAmount()
// 	statisticAvg, err := dbops.GetBasicStatisticsAvg()
// 	// 热点-产业 & 主页元素-排名-按金额排序（产业排序）
// 	hotIndu, err := dbops.GetBasicRankAmount(0, 20)
// 	// 主分类资金分布饼图
// 	basicPie, err := dbops.GetBasicPie()
// 	// 主页元素-排名-按增长率排序
// 	rankGrowth, err := dbops.GetBasicRankGrowth(0, 10)
// 	// 主页元素-解析-技术
// 	techAnalysis, err := dbops.GetBasicAnalysisTech()
// 	// 主页元素-解析-产业
// 	induAnalysis, err := dbops.GetBasicAnalysisIndu()
// 	if err != nil {
// 		sendErrorResponse(c, defs.ErrorResponse{
// 			HTTPSc: http.StatusInternalServerError,
// 			Error: defs.Err{
// 				Error:     fmt.Sprintf("es search error: %s", err),
// 				ErrorCode: elasticsearchError,
// 			},
// 		})
// 	}
// 	// 合成返回
// 	resPie := pieParser(basicPie)
// 	resCloud := cloudParser(hotIndu)
// 	fundRank := rankParser(hotIndu)

// 	result := &defs.BasicResponse{
// 		Statistic: defs.Statistic{
// 			FundingsAmount: *statisticAmount,
// 			AvgInvested:    *statisticAvg,
// 		},
// 		ResPie:   *resPie,
// 		ResCloud: *resCloud,
// 	}

// }

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

// func rankParser(hotIndu []*dbops.BasicRankAmountResult) *defs.FundRankResponse {
// 	data := make([]*defs.FundRankData, len(hotIndu))
// 	orderType = "amount"

// }

func SwitchRank(c *gin.Context) {
	return
}

func SwitchTech(c *gin.Context) {
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
