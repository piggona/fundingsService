package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/dbops"
	"github.com/piggona/fundingsView/api/defs"
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
	return
}
