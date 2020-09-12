package dbops

import (
	"encoding/json"
	"testing"

	"github.com/piggona/fundingsView/api/defs"
)

// func TestConn(t *testing.T) {
// 	ais, _ := esconn.NewAwardIDSearcher()
// 	var result interface{}
// 	award := &AwardIDResult{}
// 	result, err := ais.Request("1902389", "nsf_test")
// 	if err != nil {
// 		t.Errorf("%s", err)
// 	} else {
// 		temp, _ := json.Marshal(result)
// 		json.Unmarshal(temp, award)
// 		t.Errorf("%v", award)
// 	}
// }

func TestGetAwardID(t *testing.T) {
	result, err := GetAwardIDResult("1902389", "nsf_test")
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Errorf("%v", result)
}

func TestGetAwardIDDetail(t *testing.T) {
	result, err := GetAwardIDDetail("1902389", "nsf_test")
	if err != nil {
		t.Errorf("%s", err)
	}
	res, _ := json.Marshal(result)
	t.Errorf("%s", res)
}

func TestMultiSearch(t *testing.T) {
	amount := &defs.AwardAmount{
		Op:  0,
		Gte: 2700,
		Lte: 10000,
	}
	sort := &defs.Sort{
		Field: "Amount",
		Order: "desc",
	}
	input := &defs.MultiSearchQuery{
		Amount: []*defs.AwardAmount{amount},
		Sort:   []*defs.Sort{sort},
	}
	result, err := GetMultiSearchResult(input, "nsf_test", []string{})
	if err != nil {
		t.Errorf("%s", err)
	}
	for _, val := range result {
		t.Errorf("%v\n", val)
	}
}

func TestMultiSearchAgg(t *testing.T) {
	agg := &defs.Aggs{
		Bucket: "reference",
		Sort:   "sum",
		Order:  false,
	}
	input := &defs.MultiSearchQuery{
		Aggs: agg,
		Sort: []*defs.Sort{
			&defs.Sort{
				Field: "Amount",
				Order: "desc",
			},
		},
		All: true,
	}

	result, err := GetAggSearchResult(input, "nsf_test", []string{})
	if err != nil {
		t.Errorf("%s", err)
	}
	search := result.Search
	for _, val := range search {
		t.Errorf("%v\n", val)
	}
	aggResult := result.Agg
	for _, val := range aggResult {
		t.Errorf("%s %v\n", val.Key, val.MetricVal.Value)
	}
	t.Errorf("%v\n", result.Last)
}

func TestMultiSearchStats(t *testing.T) {
	agg := &defs.Aggs{
		Sort:  "stats",
		Order: false,
	}
	input := &defs.MultiSearchQuery{
		Aggs: agg,
		All:  true,
	}

	result, err := GetBucketSearchResult(input, "nsf_test", []string{})
	if err != nil {
		t.Errorf("%s", err)
	}
	// search := result.Search
	// for _, val := range search {
	// 	t.Errorf("%v\n", val)
	// }
	aggResult := result.Stats
	t.Errorf("%v\n", aggResult)
}

func TestFundRank(t *testing.T) {
	input := defs.MultiSearchQuery{
		All: true,
		Sort: []*defs.Sort{
			&defs.Sort{
				Field: "Amount",
				Order: "desc",
			},
		},
	}
	result, err := GetMultiSearchResult(&input, "nsf_test", []string{})
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Errorf("%v", result[0])
}

func TestFundReference(t *testing.T) {
	input := defs.MultiSearchQuery{
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
	result, err := GetAggSearchResult(&input, "nsf_test", []string{})
	if err != nil {
		t.Errorf("%v", err)
	}
	for _, val := range result.Agg {
		t.Errorf("%v\n", val)
	}
}
