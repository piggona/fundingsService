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
	techamount   = "techamount"
	techorgrank  = "techorg"
	techindurank = "techindu"
	techdivrank  = "techdiv"
)

// 技术详情页-基金投资金额排名
func GetTechnologyInvestRank(technology string) ([]*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := techamount
	params := map[string]string{
		"technology": technology,
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

// 技术详情页-机构投资排名
type TechOrgRankResult struct {
	Key       string
	DateValue map[string]int
}

type TechOrgRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type TechOrgRankYear struct {
	Buckets []*TechOrgRankYearBucket `json:"buckets"`
}

type TechOrgRankBodyElement struct {
	Key        string           `json:"key"`
	YearBucket *TechOrgRankYear `json:"year_bucket"`
}

func GetTechOrgRank(technology string) ([]*TechOrgRankResult, error) {
	bodyElement := []*TechOrgRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"technology": technology,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(techorgrank, params), INDEX)
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
	result := make([]*TechOrgRankResult, len(bodyElement))
	for id, element := range bodyElement {
		result[id] = &TechOrgRankResult{
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

// 技术详情页-产业投资排名
type TechInduRankResult struct {
	Key       string
	DateValue map[string]int
}

type TechInduRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type TechInduRankYear struct {
	Buckets []*TechInduRankYearBucket `json:"buckets"`
}

type TechInduRankBodyElement struct {
	Key        string            `json:"key"`
	YearBucket *TechInduRankYear `json:"year_bucket"`
}

func GetTechInduRank(technology string) ([]*TechInduRankResult, error) {
	bodyElement := []*TechInduRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"technology": technology,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(techindurank, params), INDEX)
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
	result := make([]*TechInduRankResult, len(bodyElement))
	for id, element := range bodyElement {
		result[id] = &TechInduRankResult{
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

// 技术详情页-主分类投资排名
type TechDivRankResult struct {
	Key       string
	Name      string
	DateValue map[string]int
}

type TechDivRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type TechDivRankYear struct {
	Buckets []*TechDivRankYearBucket `json:"buckets"`
}

type TechDivRankNameElement struct {
	Key        string           `json:"key"`
	YearBucket *TechDivRankYear `json:"year_bucket"`
}

type TechDivRankNameElementName struct {
	Buckets []*TechDivRankNameElement `json:"buckets"`
}

type TechDivRankBodyElement struct {
	Key          string                      `json:"key"`
	DivisionName *TechDivRankNameElementName `json:"division_name"`
}

func GetTechDivRank(technology string) ([]*TechDivRankResult, error) {
	bodyElement := []*TechDivRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"technology": technology,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(techdivrank, params), INDEX)
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
	result := make([]*TechDivRankResult, len(bodyElement))
	for id, element := range bodyElement {
		result[id] = &TechDivRankResult{
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
