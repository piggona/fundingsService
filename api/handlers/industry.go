package handlers

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/dbops"
	"github.com/piggona/fundingsView/api/defs"
	"github.com/piggona/fundingsView/api/utils/log"
)

func GetIndustryDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	resp := &defs.IndustryTitleDataResponse{
		Title: uuid,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

func GetIndustryRelatedFund(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetIndustryInvestRank(uuid)
	if err != nil {
		log.Error("dbops GetDivisionInvestRank error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	ranklist := make([]*defs.IndustryRankList, len(elements))
	for id, element := range elements {
		data := &defs.IndustryRankList{
			Rank:  id,
			UUID:  element.AwardID,
			Title: element.AwardTitle,
			Date:  element.AwardAmount,
		}
		ranklist[id] = data
	}

	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp: &defs.IndustryRelatedFundResponse{
			RankList: ranklist,
		},
	})
	return
}

func GetIndustryRelatedDiv(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetInduDivRank(uuid)
	if err != nil {
		log.Error("dbops GetTechDivRank error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	years := Year{}
	yearMap := map[string]struct{}{}
	for _, element := range elements {
		for date, _ := range element.DateValue {
			if _, ok := yearMap[date]; !ok {
				years = Year(append([]string(years), date))
				yearMap[date] = struct{}{}
			}
		}
	}
	sort.Sort(&years)
	rankList := make([]*defs.IndustryRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.IndustryRankList{}
		gramData[id].Name = element.Key
		data := make([]string, years.Len())
		for i, year := range []string(years) {
			value, ok := element.DateValue[year]
			if !ok {
				data[i] = "0"
				continue
			}
			data[i] = strconv.Itoa(value)
		}
		gramData[id].Data = data
		rankList[id].Title = element.Name
		rankList[id].UUID = element.Key
		rankList[id].Rank = id
		rankList[id].Date = element.DateValue[[]string(years)[years.Len()-1]]
	}
	resp := &defs.IndustryRelatedTechResponse{
		DataOne: &defs.GramData{
			XAxis: []string(years),
			Data:  gramData,
		},
		DataTwo: &defs.GramData{
			XAxis: []string(years),
			Data:  gramData,
		},
		RankList: rankList,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

func GetIndustryRelatedOrg(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetInduOrgRank(uuid)
	if err != nil {
		log.Error("dbops GetInduOrgRank error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	years := Year{}
	yearMap := map[string]struct{}{}
	for _, element := range elements {
		for date, _ := range element.DateValue {
			if _, ok := yearMap[date]; !ok {
				years = Year(append([]string(years), date))
				yearMap[date] = struct{}{}
			}
		}
	}
	sort.Sort(&years)
	rankList := make([]*defs.IndustryRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.IndustryRankList{}
		gramData[id].Name = element.Key
		data := make([]string, years.Len())
		for i, year := range []string(years) {
			value, ok := element.DateValue[year]
			if !ok {
				data[i] = "0"
				continue
			}
			data[i] = strconv.Itoa(value)
		}
		gramData[id].Data = data
		rankList[id].Title = element.Key
		rankList[id].UUID = element.Key
		rankList[id].Rank = id
		rankList[id].Date = element.DateValue[[]string(years)[years.Len()-1]]
	}
	resp := &defs.IndustryRelatedOrgResponse{
		DataOne: &defs.GramData{
			XAxis: []string(years),
			Data:  gramData,
		},
		DataTwo: &defs.GramData{
			XAxis: []string(years),
			Data:  gramData,
		},
		RankList: rankList,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

func GetIndustryRelatedTech(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetInduTechRank(uuid)
	if err != nil {
		log.Error("dbops GetInduTechRank error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	years := Year{}
	yearMap := map[string]struct{}{}
	for _, element := range elements {
		for date, _ := range element.DateValue {
			if _, ok := yearMap[date]; !ok {
				years = Year(append([]string(years), date))
				yearMap[date] = struct{}{}
			}
		}
	}
	sort.Sort(&years)
	rankList := make([]*defs.IndustryRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.IndustryRankList{}
		gramData[id].Name = element.Key
		data := make([]string, years.Len())
		for i, year := range []string(years) {
			value, ok := element.DateValue[year]
			if !ok {
				data[i] = "0"
				continue
			}
			data[i] = strconv.Itoa(value)
		}
		gramData[id].Data = data
		rankList[id].Title = element.Key
		rankList[id].UUID = element.Key
		rankList[id].Rank = id
		rankList[id].Date = element.DateValue[[]string(years)[years.Len()-1]]
	}
	resp := &defs.IndustryRelatedInduResponse{
		DataOne: &defs.GramData{
			XAxis: []string(years),
			Data:  gramData,
		},
		DataTwo: &defs.GramData{
			XAxis: []string(years),
			Data:  gramData,
		},
		RankList: rankList,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}
