package dbops

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/piggona/fundingsView/api/esconn"
	"github.com/piggona/fundingsView/api/esconn/searcher"
	"github.com/piggona/fundingsView/api/utils/log"
)

const (
	INDEX = "nsf_data"
)

// 主页元素-解析-技术
var basic_analysis_tech = `
{
  "size": 0,
  "aggs": {
    "tech_group": {
      "terms": {
        "field": "organization.division.keyword",
        "size": 100
      },
      "aggs": {
        "techs": {
          "terms": {
            "field": "technology.keyword",
            "size": 100
          },
          "aggs": {
            "tech_group_sum": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort":{
              "bucket_sort": {
                "sort": {
                  "tech_group_sum": {
                    "order": "desc"
                  }
                },
                "size": 10
              }
            }
          }
        }
      }
    }
  }
}
`

type BasicAnalysisTech struct {
	Tech  string
	Value int
}

type BasicAnalysisTechResult struct {
	Key   string
	Techs []*BasicAnalysisTech
}

type BasicAnalysisTechBodyTechElement struct {
	Key          string             `json:"key"`
	TechGroupSum map[string]float32 `json:"tech_group_sum"`
}

type BasicAnalysisTechBodyElement struct {
	Key     string                              `json:"key"`
	Buckets []*BasicAnalysisTechBodyTechElement `json:"buckets"`
}

func GetBasicAnalysisTech() ([]*BasicAnalysisTechResult, error) {
	bodyElement := []*BasicAnalysisTechBodyElement{}
	ais, err := esconn.NewAwardPlainSearcher([]string{})
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, err := ais.Request(ctx, searcher.NewPlainSearchReq(&basic_analysis_tech), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "tech_group", "buckets")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(simplejson.Json)
	if !ok {
		log.Error("assertion error when assert object to simplejson.Json: %s", err)
		return nil, err
	}
	objStr, err := jsonObj.Encode()
	if err != nil {
		log.Error("json object encode error: %s", err)
		return nil, err
	}
	err = json.Unmarshal(objStr, &bodyElement)
	if err != nil {
		log.Error("json string unmarshal to BasicAnalysisTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*BasicAnalysisTechResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		techs := make([]*BasicAnalysisTech, len(ele.Buckets))
		for i, bucket := range ele.Buckets {
			techs[i] = &BasicAnalysisTech{
				Tech:  bucket.Key,
				Value: int(bucket.TechGroupSum["value"]),
			}
		}
		result[id] = &BasicAnalysisTechResult{
			Key:   ele.Key,
			Techs: techs,
		}
	}
	return result, nil
}
