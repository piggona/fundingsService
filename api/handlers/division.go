package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/dbops"
	"github.com/piggona/fundingsView/api/defs"
	"github.com/piggona/fundingsView/api/utils/log"
)

func GetCategoryDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	resp := &defs.CategoryTitleDataResponse{
		Title: uuid,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

func GetCategoryRelatedFund(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetDivisionInvestRank(uuid)
	if err != nil {
		log.Error("dbops GetDivisionInvestRank error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	ranklist := make([]*defs.CategoryRankList, len(elements))
	for id, element := range elements {
		data := &defs.CategoryRankList{
			Rank:  id,
			UUID:  element.AwardID,
			Title: element.AwardTitle,
			Date:  element.AwardAmount,
		}
		ranklist[id] = data
	}

	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp: &defs.CategoryRelatedFundResponse{
			RankList: ranklist,
		},
	})
	return
}

type Year []string

func (y *Year) Len() int {
	return len([]string(*y))
}

func (y *Year) Less(i, j int) bool {
	yearI, _ := time.Parse((*y)[i], "01/02/2006")
	yearJ, _ := time.Parse((*y)[j], "01/02/2006")
	return yearI.Before(yearJ)
}

func (y *Year) Swap(i, j int) {
	(*y)[i], (*y)[j] = (*y)[j], (*y)[i]
}

func GetCategoryRelatedTech(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetDivTechRank(uuid)
	if err != nil {
		log.Error("dbops GetDivTechRank error: %s", err)
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
	rankList := make([]*defs.CategoryRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.CategoryRankList{}
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
	resp := &defs.CategoryRelatedTechResponse{
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

func GetCategoryRelatedOrg(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetDivOrgRank(uuid)
	if err != nil {
		log.Error("dbops GetDivOrgRank error: %s", err)
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
	rankList := make([]*defs.CategoryRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.CategoryRankList{}
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
	resp := &defs.CategoryRelatedOrgResponse{
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

func GetCategoryRelatedIndu(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetDivInduRank(uuid)
	if err != nil {
		log.Error("dbops GetDivInduRank error: %s", err)
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
	rankList := make([]*defs.CategoryRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.CategoryRankList{}
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
	resp := &defs.CategoryRelatedInduResponse{
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
