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
	induamount   = "induamount"
	induorgrank  = "induorg"
	indutechrank = "indutech"
	indudivrank  = "indudiv"
)

// 产业详情页-基金投资金额排名
func GetIndustryInvestRank(industry string) ([]*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := induamount
	params := map[string]string{
		"industry": industry,
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

// 产业详情页-机构投资排名
type InduOrgRankResult struct {
	Key       string
	DateValue map[string]int
}

type InduOrgRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type InduOrgRankYear struct {
	Buckets []*InduOrgRankYearBucket `json:"buckets"`
}

type InduOrgRankBodyElement struct {
	Key        string           `json:"key"`
	YearBucket *InduOrgRankYear `json:"year_bucket"`
}

func GetInduOrgRank(industry string) ([]*InduOrgRankResult, error) {
	bodyElement := []*InduOrgRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"industry": industry,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(induorgrank, params), INDEX)
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
	result := make([]*InduOrgRankResult, len(bodyElement))
	for id, element := range bodyElement {
		result[id] = &InduOrgRankResult{
			DateValue: map[string]int{},
		}
		ele := element
		for _, buc := range ele.YearBucket.Buckets {
			bucket := buc
			result[id].DateValue[bucket.Date] = int(bucket.DateValue["value"])
		}
		result[id].Key = element.Key
	}
	return result, nil
}

// 产业详情页-技术投资排名
type InduTechRankResult struct {
	Key       string
	DateValue map[string]int
}

type InduTechRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type InduTechRankYear struct {
	Buckets []*InduTechRankYearBucket `json:"buckets"`
}

type InduTechRankBodyElement struct {
	Key        string            `json:"key"`
	YearBucket *InduTechRankYear `json:"year_bucket"`
}

func GetInduTechRank(industry string) ([]*InduTechRankResult, error) {
	bodyElement := []*InduTechRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"industry": industry,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(indutechrank, params), INDEX)
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
	result := make([]*InduTechRankResult, len(bodyElement))
	for id, element := range bodyElement {
		result[id] = &InduTechRankResult{
			DateValue: map[string]int{},
		}
		ele := element
		for _, buc := range ele.YearBucket.Buckets {
			bucket := buc
			result[id].DateValue[bucket.Date] = int(bucket.DateValue["value"])
		}
		result[id].Key = element.Key
	}
	return result, nil
}

// 产业详情页-主分类投资排名
type InduDivRankResult struct {
	Key       string
	Name      string
	DateValue map[string]int
}

type InduDivRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type InduDivRankYear struct {
	Buckets []*InduDivRankYearBucket `json:"buckets"`
}

type InduDivRankNameElement struct {
	Key        string           `json:"key"`
	YearBucket *InduDivRankYear `json:"year_bucket"`
}

type InduDivRankNameElementName struct {
	Buckets []*InduDivRankNameElement `json:"buckets"`
}

type InduDivRankBodyElement struct {
	Key          string                      `json:"key"`
	DivisionName *InduDivRankNameElementName `json:"division_name"`
}

func GetInduDivRank(industry string) ([]*InduDivRankResult, error) {
	bodyElement := []*InduDivRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"industry": industry,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(indudivrank, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "division_list", "buckets")
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
	result := make([]*InduDivRankResult, len(bodyElement))
	for id, element := range bodyElement {
		result[id] = &InduDivRankResult{
			DateValue: map[string]int{},
		}
		ele := element
		for _, buc := range ele.DivisionName.Buckets[0].YearBucket.Buckets {
			bucket := buc
			result[id].DateValue[bucket.Date] = int(bucket.DateValue["value"])
		}
		result[id].Key = element.Key
		result[id].Name = element.DivisionName.Buckets[0].Key
	}
	return result, nil
}
