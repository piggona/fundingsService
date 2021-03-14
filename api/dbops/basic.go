package dbops

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/piggona/fundingsView/api/esconn"
	"github.com/piggona/fundingsView/api/esconn/searcher"
	"github.com/piggona/fundingsView/api/utils/log"
)

const (
	INDEX = "nsf_data"

	// template ids
	basicrankgrowth = "basicrankgrowth"
	basicrankamount = "basicrankamount"
)

// 主页元素-排名-按金额排序(产业排序)
type BasicRankAmountResult struct {
	Key   string
	Value int
}

type BasicRankAmountBodyElement struct {
	Key       string             `json:"key"`
	InduValue map[string]float64 `json:"category_proportion"`
}

func GetBasicRankAmount(from, size int) ([]*BasicRankAmountResult, error) {
	bodyElement := []*BasicRankAmountBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := basicrankamount
	params := map[string]string{
		"from": strconv.Itoa(from),
		"size": strconv.Itoa(size),
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(templateId, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "amount_list", "buckets")
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
	fmt.Println(string(objStr))
	err = json.Unmarshal(objStr, &bodyElement)
	if err != nil {
		log.Error("json string unmarshal to BasicAnalysisTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*BasicRankAmountResult, len(bodyElement))
	for id, bucs := range bodyElement {
		buckets := bucs
		result[id] = &BasicRankAmountResult{
			Key:   buckets.Key,
			Value: int(buckets.InduValue["value"]),
		}
	}
	return result, nil
}

// 主页元素-排名-按增长率排序
type BasicRankGrowthResult struct {
	Key       string
	DateValue map[string]int
}

type BasicRankGrowthBodyBucket struct {
	KeyAsString string             `json:"key_as_string"`
	RateSum     map[string]float64 `json:"rate_sum"`
}

type AggYear struct {
	Buckets []*BasicRankGrowthBodyBucket `json:"buckets"`
}

type BasicRankGrowthBodyElement struct {
	Key     string   `json:"key"`
	AggYear *AggYear `json:"agg_year"`
}

func GetBasicRankGrowth(from, size int) ([]*BasicRankGrowthResult, error) {
	bodyElement := []*BasicRankGrowthBodyElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := basicrankgrowth
	params := map[string]string{
		"from": strconv.Itoa(from),
		"size": strconv.Itoa(size),
	}
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(templateId, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "rate_list", "buckets")
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
		log.Error("json string unmarshal to BasicAnalysisTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*BasicRankGrowthResult, len(bodyElement))
	for id, bucs := range bodyElement {
		buckets := bucs
		result[id] = &BasicRankGrowthResult{
			Key:       buckets.Key,
			DateValue: make(map[string]int),
		}
		for _, buc := range buckets.AggYear.Buckets {
			bucket := buc
			result[id].DateValue[bucket.KeyAsString] = int(bucket.RateSum["value"])
		}
	}
	return result, nil
}

// 主页元素-解析-技术
type BasicAnalysisTech struct {
	Tech  string `json:"tech"`
	Value int    `json:"amount"`
}

type BasicAnalysisTechResult struct {
	Key   string
	Name  string
	Techs []*BasicAnalysisTech
}

type BasicAnalysisTechBodyTechElement struct {
	Key          string             `json:"key"`
	TechGroupSum map[string]float32 `json:"tech_group_sum"`
}

type Techs struct {
	Buckets []*BasicAnalysisTechBodyTechElement `json:"buckets"`
}

type DivisionNameBucket struct {
	Key   string `json:"key"`
	Techs Techs  `json:"techs"`
}

type DivisionName struct {
	Buckets []*DivisionNameBucket `json:"buckets"`
}

type BasicAnalysisTechBodyElement struct {
	Key          string       `json:"key"`
	DivisionName DivisionName `json:"division_name"`
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
		log.Error("json string unmarshal to BasicAnalysisTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*BasicAnalysisTechResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		techs := make([]*BasicAnalysisTech, len(ele.DivisionName.Buckets[0].Techs.Buckets))
		for i, buc := range ele.DivisionName.Buckets[0].Techs.Buckets {
			bucket := buc
			techs[i] = &BasicAnalysisTech{
				Tech:  bucket.Key,
				Value: int(bucket.TechGroupSum["value"]),
			}
		}
		result[id] = &BasicAnalysisTechResult{
			Key:   ele.Key,
			Name:  ele.DivisionName.Buckets[0].Key,
			Techs: techs,
		}
	}
	return result, nil
}

// 主页元素-解析-产业
type BasicAnalysisIndu struct {
	Indu  string `json:"tech"`
	Value int    `json:"amount"`
}

type BasicAnalysisInduResult struct {
	Key   string
	Name  string
	Indus []*BasicAnalysisIndu
}

type BasicAnalysisInduBodyInduElement struct {
	Key          string             `json:"key"`
	InduGroupSum map[string]float32 `json:"indu_group_sum"`
}

type Indus struct {
	Buckets []*BasicAnalysisInduBodyInduElement `json:"buckets"`
}

type InduDivisionNameBucket struct {
	Key   string `json:"key"`
	Indus Indus  `json:"indus"`
}

type InduDivisionName struct {
	Buckets []*InduDivisionNameBucket `json:"buckets"`
}

type BasicAnalysisInduBodyElement struct {
	Key          string           `json:"key"`
	DivisionName InduDivisionName `json:"division_name"`
}

func GetBasicAnalysisIndu() ([]*BasicAnalysisInduResult, error) {
	bodyElement := []*BasicAnalysisInduBodyElement{}
	ais, err := esconn.NewAwardPlainSearcher([]string{})
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, err := ais.Request(ctx, searcher.NewPlainSearchReq(&basic_analysis_indu), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "indu_group", "buckets")
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
		log.Error("json string unmarshal to BasicAnalysisInduBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*BasicAnalysisInduResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		indus := make([]*BasicAnalysisIndu, len(ele.DivisionName.Buckets[0].Indus.Buckets))
		for i, buc := range ele.DivisionName.Buckets[0].Indus.Buckets {
			bucket := buc
			indus[i] = &BasicAnalysisIndu{
				Indu:  bucket.Key,
				Value: int(bucket.InduGroupSum["value"]),
			}
		}
		result[id] = &BasicAnalysisInduResult{
			Key:   ele.Key,
			Name:  ele.DivisionName.Buckets[0].Key,
			Indus: indus,
		}
	}
	return result, nil
}

// 主页元素-统计-基金总投资数
func GetBasicStatisticsAmount() (*float64, error) {
	ais, err := esconn.NewAwardPlainSearcher([]string{})
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, err := ais.Request(ctx, searcher.NewPlainSearchReq(&basic_statistics_amount), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "fundings_amount", "value")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
	if !ok {
		log.Error("assertion error when assert object to simplejson.Json: %s", err)
		return nil, err
	}
	result, err := jsonObj.Float64()
	if err != nil {
		log.Error("json convert to float64 error: %s", err)
		return nil, err
	}
	return &result, nil
}

// 主页元素-统计-基金平均投资数
func GetBasicStatisticsAvg() (*float64, error) {
	ais, err := esconn.NewAwardPlainSearcher([]string{})
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, err := ais.Request(ctx, searcher.NewPlainSearchReq(&basic_statistics_avg), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "avg_invested", "value")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
	if !ok {
		log.Error("assertion error when assert object to simplejson.Json: %s", err)
		return nil, err
	}
	result, err := jsonObj.Float64()
	if err != nil {
		log.Error("json convert to float64 error: %s", err)
		return nil, err
	}
	return &result, nil
}

// 热点产业(使用：主页元素-排名-按金额排序，form:0,size:20)

// 主分类资金分布饼图
type BasicPie struct {
	Key   string
	Name  string
	Value int
}

type BasicCategoryNameBucket struct {
	Key       string             `json:"key"`
	CateValue map[string]float64 `json:"category_proportion"`
}

type BasicCategoryName struct {
	Buckets []*BasicCategoryNameBucket `json:"buckets"`
}

type BasicCategoriesBodyElement struct {
	Key          string             `json:"key"`
	CategoryName *BasicCategoryName `json:"category_name"`
}

func GetBasicPie() ([]*BasicPie, error) {
	bodyElement := []*BasicCategoriesBodyElement{}
	ais, err := esconn.NewAwardPlainSearcher([]string{})
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, err := ais.Request(ctx, searcher.NewPlainSearchReq(&basic_pie), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "categorys", "buckets")
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
		log.Error("json string unmarshal to BasicCategoriesBodyElements error: %s", err)
		return nil, err
	}
	result := make([]*BasicPie, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &BasicPie{
			Key:   ele.Key,
			Name:  ele.CategoryName.Buckets[0].Key,
			Value: int(ele.CategoryName.Buckets[0].CateValue["value"]),
		}
	}
	return result, nil
}
