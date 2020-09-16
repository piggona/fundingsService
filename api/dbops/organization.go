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
	orgamount   = "orgamount"
	orgtechrank = "orgtech"
	orgindurank = "orgindu"
	orgdivrank  = "orgdiv"
)

// 机构详情页-基金投资金额排名
func GetOrganizationInvestRank(organization string) ([]*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := orgamount
	params := map[string]string{
		"organization": organization,
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

// 机构详情页-技术投资排名
type OrgTechRankResult struct {
	Key       string
	DateValue map[string]int
}

type OrgTechRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type OrgTechRankYear struct {
	Buckets []*OrgTechRankYearBucket `json:"buckets"`
}

type OrgTechRankBodyElement struct {
	Key        string           `json:"key"`
	YearBucket *OrgTechRankYear `json:"year_bucket"`
}

func GetOrgTechRank(organization string) ([]*OrgTechRankResult, error) {
	bodyElement := []*OrgTechRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"organization": organization,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(orgtechrank, params), INDEX)
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
	result := make([]*OrgTechRankResult, len(bodyElement))
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

// 机构详情页-产业投资排名
type OrgInduRankResult struct {
	Key       string
	DateValue map[string]int
}

type OrgInduRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type OrgInduRankYear struct {
	Buckets []*OrgInduRankYearBucket `json:"buckets"`
}

type OrgInduRankBodyElement struct {
	Key        string           `json:"key"`
	YearBucket *OrgInduRankYear `json:"year_bucket"`
}

func GetOrgInduRank(organization string) ([]*OrgInduRankResult, error) {
	bodyElement := []*OrgInduRankBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	params := map[string]string{
		"organization": organization,
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(orgindurank, params), INDEX)
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
	result := make([]*OrgInduRankResult, len(bodyElement))
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

// 机构详情页-主分类投资排名
type OrgDivRankResult struct {
	Key       string
	Name      string
	DateValue map[string]int
}

type OrgDivRankYearBucket struct {
	Date      string             `json:"key_as_string"`
	DateValue map[string]float64 `json:"proportion_fund"`
}

type OrgDivRankYear struct {
	Buckets []*OrgDivRankYearBucket `json:"buckets"`
}

type OrgDivRankNameElement struct {
	Key        string          `json:"key"`
	YearBucket *OrgDivRankYear `json:"year_bucket"`
}

type OrgDivRankNameElementName struct {
	Buckets []*OrgDivRankNameElement `json:"buckets"`
}

type OrgDivRankBodyElement struct {
	Key          string                     `json:"key"`
	DivisionName *OrgDivRankNameElementName `json:"division_name"`
}

func GetOrgDivRank(division string) ([]*OrgDivRankResult, error) {
	bodyElement := []*OrgDivRankBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(orgdivrank, params), INDEX)
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
	result := make([]*OrgDivRankResult, len(bodyElement))
	for id, element := range bodyElement {
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
