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

func HomePage(c *gin.Context) {
	result := &defs.BasicResponse{}
	// statistic部分：总投资金额&平均投资金额
	agg := &defs.Aggs{
		Sort:  "stats",
		Order: false,
	}
	totalAmount := &defs.MultiSearchQuery{
		All:  true,
		Aggs: agg,
	}
	restotal, err := dbops.GetBucketSearchResult(totalAmount, "nsf_test", []string{})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	result.Statistic.FundingsAmount = restotal.Stats.Sum
	result.Statistic.AvgInvested = restotal.Stats.Avg
	// FundRank部分：按金额排序&按增长率排序
	amountRank := &defs.MultiSearchQuery{
		All: true,
		Sort: []*defs.Sort{
			&defs.Sort{
				Field: "Amount",
				Order: "desc",
			},
		},
	}
	resAmountRank, err := dbops.GetMultiSearchResult(amountRank, "nsf_test", []string{})
	result.FundRank.OrderType = "desc"
	for id, val := range resAmountRank {
		temp := &defs.FundRankData{
			Rank:  strconv.Itoa(id),
			Name:  val.AwardTitle,
			Money: val.AwardAmount,
		}
		result.FundRank.Data = append(result.FundRank.Data, temp)
	}
	// PieGraph部分：按金额顺序排序主分类
	reference := &defs.MultiSearchQuery{
		All: true,
		Sort: []*defs.Sort{
			&defs.Sort{
				Field: "Amount",
				Order: "desc",
			},
		},
		Aggs: &defs.Aggs{
			Bucket: "reference",
			Sort:   "sum",
			Order:  false,
		},
	}
	resRef, err := dbops.GetAggSearchResult(reference, "nsf_test", []string{})
	for _, val := range resRef.Agg {
		result.ResPie.Legend = append(result.ResPie.Legend, val.Key)
		result.ResPie.Series = append(result.ResPie.Series, &defs.SeriesData{Value: strconv.Itoa(val.MetricVal.Value), Name: val.Key})
	}
	c.JSON(http.StatusOK, result)
}

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
	dbops.GetMultiSearch(dbSearch)
	return
}
