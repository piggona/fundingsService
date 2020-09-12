package searcher

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/piggona/fundingsView/api/defs"
)

type AwardAggSearch struct {
	AwardMultiSearch
}

func (a *AwardAggSearch) parser(input SearcherReq) (map[string]interface{}, error) {
	var result map[string]interface{}
	tm, ok := input.(*AggSearcherReq)
	if !ok {
		return nil, fmt.Errorf("Error bad request type, expect type: *AggSearcherReq")
	}
	multi := MultiSearcherReq(*tm)
	result, err := a.AwardMultiSearch.parser(&multi)
	if err != nil {
		return nil, err
	}
	ti, err := tm.Validate()
	if err != nil {
		return nil, err
	}
	i := ti.(defs.MultiSearchQuery)
	// 配置aggs
	agg := i.Aggs
	if agg != nil {
		aggtemp := map[string]interface{}{}
		var agg_not_available bool
		if len(agg.Bucket) != 0 {
			var sort string
			var order string
			var field string
			switch agg.Sort {
			case "sum":
				sort = "nsf_order.value"
			case "count":
				sort = "_count"
			default:
				sort = "default"
				agg_not_available = true
			}
			if agg.Order {
				order = "asc"
			} else {
				order = "desc"
			}
			switch agg.Bucket {
			case "institution":
				field = "doc.Award.Institution.Name.keyword"
			case "reference":
				field = "doc.Award.ProgramReference.Text.keyword"
			default:
				field = ""
			}
			aggtemp["nsf_aggs"] = map[string]interface{}{
				"terms": map[string]interface{}{
					"field": field,
					"size":  100,
					"order": map[string]interface{}{
						sort: order,
					},
				},
				"aggs": map[string]interface{}{
					"nsf_order": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "doc.Award.AwardAmount",
						},
					},
				},
			}
		} else {
			if agg.Sort == "stats" {
				aggtemp["nsf_aggs"] = map[string]interface{}{
					"stats": map[string]interface{}{
						"field": "doc.Award.AwardAmount",
					},
				}
			} else {
				agg_not_available = true
			}
		}
		if !agg_not_available {
			result["aggs"] = aggtemp
		}
	}
	return result, nil
}

func (a *AwardAggSearch) Request(input SearcherReq, index string) (SearcherResp, error) {
	searchmap, err := a.parser(input)
	if err != nil {
		return nil, err
	}
	res, err := a.conn.Query(searchmap, index)
	if err != nil {
		return nil, err
	}
	result, err := NewAggSearcherResp(&res)
	// fmt.Printf("debug!!! %v\n", result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AggSearcherReq defs.MultiSearchQuery

func (i *AggSearcherReq) Validate() (interface{}, error) {
	return defs.MultiSearchQuery(*i), nil
}

func NewAggSearcherReq(req *MultiSearchQuery) (*AggSearcherReq, error) {
	id := AggSearcherReq(*req)
	return &id, nil
}

type AggSearcherResp map[string]interface{}

func (idsearch *AggSearcherResp) Find(resp interface{}, paths ...string) (interface{}, error) {
	var temp interface{}
	var path string
	temp = resp
	for i := 0; i < len(paths); i++ {
		path = paths[i]
		t, err := strconv.Atoi(path)
		if err != nil {
			// 不是数字
			if path == "[]" {
				var tres []interface{}
				tt, ok := temp.([]interface{})
				if !ok {
					return nil, errors.New("Error in type assertion []interface{}")
				}
				for _, val := range tt {
					item, err := idsearch.Find(val, paths[i+1:]...)
					if err != nil {
						return nil, err
					}
					tres = append(tres, item)
				}
				return tres, nil
			}
			tt, ok := temp.(*AggSearcherResp)
			if !ok {
				tnormal, ok := temp.(map[string]interface{})
				if !ok {
					return nil, errors.New("Error in type assertion map[string]interface{}")
				}
				temp = tnormal[path]
			} else {
				temp = (*tt)[path]
			}
		} else {
			// 数字
			tt, ok := temp.([]interface{})
			if !ok {
				return nil, errors.New("Error in type assertion []interface{}")
			}
			temp = tt[t]
		}
	}
	return temp, nil
}

func NewAggSearcherResp(resp interface{}) (*AggSearcherResp, error) {
	res, ok := resp.(*map[string]interface{})
	if !ok {
		return nil, errors.New("Error bad type resp.")
	}
	result := AggSearcherResp(*res)
	return &result, nil
}
