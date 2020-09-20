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
	funddetail   = "funddetail"
	fundindudiv  = "fundindudiv"
	fundinduorg  = "fundinduorg"
	fundindutech = "fundindutech"
	fundtechdiv  = "fundtechdiv"
	fundtechorg  = "fundtechorg"
	fundtechindu = "fundtechindu"
	funddivindu  = "funddivindu"
	funddivtech  = "funddivtech"
	funddivorg   = "funddivorg"
	fundorgtech  = "fundorgtech"
	fundorgindu  = "fundorgindu"
	fundorgdiv   = "fundorgdiv"
)

// 基金详情页-详情
func GetFundDetail(fundID string) (*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := funddetail
	params := map[string]string{
		"uuid": fundID,
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
	if len(bodyElement) == 0 {
		return nil, nil
	}
	result := bodyElement[0].Source
	return result, nil
}

// 基金详情页-关联关系
// 基金详情页-产业相关-主分类
type FundInduDivResult struct {
	Key   string
	Name  string
	Value int
}

type FundInduDivCategoryNameBucket struct {
	Key       string             `json:"key"`
	CateValue map[string]float64 `json:"proportion_fund"`
}

type FundInduDivCategoryName struct {
	Buckets []*FundInduDivCategoryNameBucket `json:"buckets"`
}

type FundInduDivBodyElement struct {
	Key          string                   `json:"key"`
	CategoryName *FundInduDivCategoryName `json:"division_name"`
}

func GetFundInduDiv(industry string) ([]*FundInduDivResult, error) {
	bodyElement := []*FundInduDivBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundindudiv, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_division", "buckets")
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
	result := make([]*FundInduDivResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundInduDivResult{
			Key:   ele.Key,
			Name:  ele.CategoryName.Buckets[0].Key,
			Value: int(ele.CategoryName.Buckets[0].CateValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-产业相关-机构
type FundInduOrgResult struct {
	Key   string
	Value int
}

type FundInduOrgBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundInduOrg(industry string) ([]*FundInduOrgResult, error) {
	bodyElement := []*FundInduOrgBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundinduorg, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_organizations", "buckets")
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
		log.Error("json string unmarshal to FundInduOrgBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundInduOrgResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundInduOrgResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-产业相关-技术
type FundInduTechResult struct {
	Key   string
	Value int
}

type FundInduTechBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundInduTech(industry string) ([]*FundInduTechResult, error) {
	bodyElement := []*FundInduTechBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundindutech, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_technologies", "buckets")
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
		log.Error("json string unmarshal to FundInduTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundInduTechResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundInduTechResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-技术相关-主分类
type FundTechDivResult struct {
	Key   string
	Name  string
	Value int
}

type FundTechDivCategoryNameBucket struct {
	Key       string             `json:"key"`
	CateValue map[string]float64 `json:"proportion_fund"`
}

type FundTechDivCategoryName struct {
	Buckets []*FundTechDivCategoryNameBucket `json:"buckets"`
}

type FundTechDivBodyElement struct {
	Key          string                   `json:"key"`
	CategoryName *FundTechDivCategoryName `json:"division_name"`
}

func GetFundTechDiv(technology string) ([]*FundTechDivResult, error) {
	bodyElement := []*FundTechDivBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundtechdiv, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_division", "buckets")
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
		log.Error("json string unmarshal to FundTechDivBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundTechDivResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundTechDivResult{
			Key:   ele.Key,
			Name:  ele.CategoryName.Buckets[0].Key,
			Value: int(ele.CategoryName.Buckets[0].CateValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-技术相关-产业
type FundTechInduResult struct {
	Key   string
	Value int
}

type FundTechInduBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundTechIndu(technology string) ([]*FundTechInduResult, error) {
	bodyElement := []*FundTechInduBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundtechindu, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_industries", "buckets")
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
		log.Error("json string unmarshal to FundTechInduBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundTechInduResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundTechInduResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-技术相关-机构
type FundTechOrgResult struct {
	Key   string
	Value int
}

type FundTechOrgBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundTechOrg(technology string) ([]*FundTechOrgResult, error) {
	bodyElement := []*FundTechOrgBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundtechorg, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_organizations", "buckets")
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
		log.Error("json string unmarshal to FundTechOrgBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundTechOrgResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundTechOrgResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-主分类相关-产业
type FundDivInduResult struct {
	Key   string
	Value int
}

type FundDivInduBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundDivIndu(division string) ([]*FundDivInduResult, error) {
	bodyElement := []*FundDivInduBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(funddivindu, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_industries", "buckets")
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
		log.Error("json string unmarshal to FundDivInduBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundDivInduResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundDivInduResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-主分类相关-技术
type FundDivTechResult struct {
	Key   string
	Value int
}

type FundDivTechBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundDivTech(division string) ([]*FundDivTechResult, error) {
	bodyElement := []*FundDivTechBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(funddivtech, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_technologies", "buckets")
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
		log.Error("json string unmarshal to FundDivTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundDivTechResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundDivTechResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-主分类相关-机构
type FundDivOrgResult struct {
	Key   string
	Value int
}

type FundDivOrgBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundDivOrg(division string) ([]*FundDivOrgResult, error) {
	bodyElement := []*FundDivOrgBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(funddivorg, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_organizations", "buckets")
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
		log.Error("json string unmarshal to FundDivOrgBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundDivOrgResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundDivOrgResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-机构相关-产业
type FundOrgInduResult struct {
	Key   string
	Value int
}

type FundOrgInduBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundOrgIndu(organization string) ([]*FundOrgInduResult, error) {
	bodyElement := []*FundOrgInduBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundorgindu, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_industries", "buckets")
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
		log.Error("json string unmarshal to FundOrgInduBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundOrgInduResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundOrgInduResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-机构相关-技术
type FundOrgTechResult struct {
	Key   string
	Value int
}

type FundOrgTechBodyElement struct {
	Key      string             `json:"key"`
	OrgValue map[string]float64 `json:"proportion_fund"`
}

func GetFundOrgTech(organization string) ([]*FundOrgTechResult, error) {
	bodyElement := []*FundOrgTechBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundorgtech, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_technologies", "buckets")
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
		log.Error("json string unmarshal to FundOrgTechBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundOrgTechResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundOrgTechResult{
			Key:   ele.Key,
			Value: int(ele.OrgValue["value"]),
		}
	}
	return result, nil
}

// 基金详情页-机构相关-主分类
type FundOrgDivResult struct {
	Key   string
	Name  string
	Value int
}

type FundOrgDivCategoryNameBucket struct {
	Key       string             `json:"key"`
	CateValue map[string]float64 `json:"proportion_fund"`
}

type FundOrgDivCategoryName struct {
	Buckets []*FundOrgDivCategoryNameBucket `json:"buckets"`
}

type FundOrgDivBodyElement struct {
	Key          string                  `json:"key"`
	CategoryName *FundOrgDivCategoryName `json:"division_name"`
}

func GetFundOrgDiv(organization string) ([]*FundOrgDivResult, error) {
	bodyElement := []*FundOrgDivBodyElement{}
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
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(fundorgdiv, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "aggregations", "related_division", "buckets")
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
		log.Error("json string unmarshal to FundOrgDivBodyElement error: %s", err)
		return nil, err
	}
	result := make([]*FundOrgDivResult, len(bodyElement))
	for id, element := range bodyElement {
		ele := element
		result[id] = &FundOrgDivResult{
			Key:   ele.Key,
			Name:  ele.CategoryName.Buckets[0].Key,
			Value: int(ele.CategoryName.Buckets[0].CateValue["value"]),
		}
	}
	return result, nil
}

func GetRelatedFunds(key, value string, page int) ([]*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardPlainSearcher([]string{})
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	searchContent := GetRelatedTemplate(key, value, page)
	resp, err := ais.Request(ctx, searcher.NewPlainSearchReq(&searchContent), INDEX)
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
