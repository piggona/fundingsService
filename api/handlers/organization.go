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

func GetOrganizationDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	resp := &defs.OrganizationTitleDataResponse{
		Title: uuid,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

func GetOrganizationRelatedFund(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetOrganizationInvestRank(uuid)
	if err != nil {
		log.Error("dbops GetOranizationInvestRank error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	ranklist := make([]*defs.OrganizationRankList, len(elements))
	for id, element := range elements {
		data := &defs.OrganizationRankList{
			Rank:  id,
			UUID:  element.AwardID,
			Title: element.AwardTitle,
			Date:  element.AwardAmount,
		}
		ranklist[id] = data
	}

	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp: &defs.OrganizationRelatedFundResponse{
			RankList: ranklist,
		},
	})
	return
}

func GetOrganizationRelatedTech(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetOrgTechRank(uuid)
	if err != nil {
		log.Error("dbops GetOrgTechRank error: %s", err)
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
	rankList := make([]*defs.OrganizationRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.OrganizationRankList{}
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
	resp := &defs.OrganizationRelatedTechResponse{
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

func GetOrganizationRelatedDiv(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetOrgDivRank(uuid)
	if err != nil {
		log.Error("dbops GetOrgOrgRank error: %s", err)
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
	rankList := make([]*defs.OrganizationRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.OrganizationRankList{}
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
	resp := &defs.OrganizationRelatedOrgResponse{
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

func GetOrganizationRelatedIndu(c *gin.Context) {
	uuid := c.Param("uuid")
	elements, err := dbops.GetOrgInduRank(uuid)
	if err != nil {
		log.Error("dbops GetOrgInduRank error: %s", err)
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
	rankList := make([]*defs.OrganizationRankList, len(elements))
	gramData := make([]*defs.GramDataBucket, len(elements))
	for id, element := range elements {
		gramData[id] = &defs.GramDataBucket{}
		rankList[id] = &defs.OrganizationRankList{}
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
	resp := &defs.OrganizationRelatedInduResponse{
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
