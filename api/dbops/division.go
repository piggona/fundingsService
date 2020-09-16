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
	divamount   = "divamount"
	divtechrank = "divtech"
	divindurank = "divindu"
	divorgrank  = "divorg"
)

// 主分类详情页-基金投资金额排名
func GetDivisionInvestRank(division string) ([]*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := divamount
	params := map[string]string{
		"division": division,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(templateId, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "hits", "hits")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
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
		log.Error("json string unmarshal to SearchResultElement error: %s", err)
		return nil, err
	}
	result := make([]*FundElement, len(bodyElement))
	for id, bucs := range bodyElement {
		buckets := bucs
		result[id] = buckets.Source
	}
	return result, nil
}

// 主分类详情页-技术投资排名
type DivTechRankResult struct {
	Key       string
	DateValue map[string]int
}

type DivTechRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type DivTechRankYear struct {
	Buckets []*DivTechRankYearBucket `json:"buckets"`
}

type DivTechRankBodyElement struct {
	Key        string           `json:"key"`
	YearBucket *DivTechRankYear `json:"year_bucket"`
}

func GetDivTechRank(division string) ([]*DivTechRankResult, error) {
	bodyElement := []*DivTechRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"division": division,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(divtechrank, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "tech_list", "buckets")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
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
		log.Error("json string unmarshal to FundInduDivBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*DivTechRankResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		for _, buc := range ele.YearBucket.Buckets {
			bucket := buc
			result[id].DateValue[bucket.Date] = int(bucket.DateValue["value"])
		}
		result[id].Key = element.Key
	}
	return result, nil
}

// 主分类详情页-产业投资排名
type DivInduRankResult struct {
	Key       string
	DateValue map[string]int
}

type DivInduRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type DivInduRankYear struct {
	Buckets []*DivInduRankYearBucket `json:"buckets"`
}

type DivInduRankBodyElement struct {
	Key        string           `json:"key"`
	YearBucket *DivInduRankYear `json:"year_bucket"`
}

func GetDivInduRank(division string) ([]*DivInduRankResult, error) {
	bodyElement := []*DivInduRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"division": division,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(divtechrank, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "indu_list", "buckets")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
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
		log.Error("json string unmarshal to FundInduDivBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*DivInduRankResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		for _, buc := range ele.YearBucket.Buckets {
			bucket := buc
			result[id].DateValue[bucket.Date] = int(bucket.DateValue["value"])
		}
		result[id].Key = element.Key
	}
	return result, nil
}

// 主分类详情页-机构投资排名
type DivOrgRankResult struct {
	Key       string
	DateValue map[string]int
}

type DivOrgRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type DivOrgRankYear struct {
	Buckets []*DivOrgRankYearBucket `json:"buckets"`
}

type DivOrgRankBodyElement struct {
	Key        string          `json:"key"`
	YearBucket *DivOrgRankYear `json:"year_bucket"`
}

func GetDivOrgRank(division string) ([]*DivOrgRankResult, error) {
	bodyElement := []*DivOrgRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"division": division,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(divorgrank, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "org_list", "buckets")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
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
		log.Error("json string unmarshal to FundInduDivBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*DivOrgRankResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		for _, buc := range ele.YearBucket.Buckets {
			bucket := buc
			result[id].DateValue[bucket.Date] = int(bucket.DateValue["value"])
		}
		result[id].Key = element.Key
	}
	return result, nil
}
