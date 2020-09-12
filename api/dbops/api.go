package dbops

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/piggona/fundingsView/api/defs"
	"github.com/piggona/fundingsView/api/esconn"
	"github.com/piggona/fundingsView/api/esconn/searcher"
)

// Part AwardIDResult

// AwardIDResult Expected AwardID search result.
type AwardIDResult struct {
	AwardID             string `json:"AwardID"`
	AwardTitle          string `json:"AwardTitle"`
	AwardAmount         string `json:"AwardAmount"`
	AwardExpirationDate string `json:"AwardExpirationDate"`
}

// GetAwardIDResult Get result from elasticsearch,and convert from map to AwardIDResult
func GetAwardIDResult(id string, index string) (*AwardIDResult, error) {
	var result interface{}
	award := &AwardIDResult{}
	ais, err := esconn.NewAwardIDSearcher()
	if err != nil {
		return nil, err
	}
	req, err := searcher.NewIDSearcherReq(id)
	if err != nil {
		return nil, errors.New("Errors Bad type 'id' not type string")
	}
	result, err = ais.Request(req, index)
	if err != nil {
		return nil, err
	}
	temp, err := json.Marshal(&result)
	if err != nil {
		return nil, fmt.Errorf("Error in json encoding: %s", err)
	}
	err = json.Unmarshal(temp, award)
	if err != nil {
		return nil, fmt.Errorf("Error in json decoding: %s", err)
	}
	return award, nil
}

type (
	BasicDetailResponse   = defs.BasicDetailResponse
	BasicOrganization     = defs.BasicOrganization
	BasicInstitutionData  = defs.BasicInstitutionData
	BasicInvestigatorData = defs.BasicInvestigatorData
	BasicProgramReference = defs.BasicProgramReference
)

func GetAwardIDDetail(id string, index string) (*BasicDetailResponse, error) {
	var result searcher.SearcherResp
	award := &BasicDetailResponse{}
	ais, err := esconn.NewAwardIDSearcher()
	if err != nil {
		return nil, err
	}
	ais.SetSource([]string{})
	req, err := searcher.NewIDSearcherReq(id)
	if err != nil {
		return nil, errors.New("Errors Bad type 'id' not type string")
	}
	result, err = ais.Request(req, index)
	if err != nil {
		return nil, err
	}

	ttitle, err := result.Find(result, []string{"AwardTitle"}...)
	if err != nil {
		return nil, err
	}
	title, ok := ttitle.(string)
	if !ok {
		return nil, fmt.Errorf("Errors Bad type 'AwardTitle' not string")
	}
	award.AwardTitle = title

	tabstract, err := result.Find(result, []string{"AbstractNarration"}...)
	if err != nil {
		return nil, err
	}
	abstract, ok := tabstract.(string)
	if !ok {
		return nil, fmt.Errorf("Errors Bad type 'AbstractNarration' not string")
	}
	award.AbstractNarration = abstract

	tawardAmount, err := result.Find(result, []string{"AwardAmount"}...)
	if err != nil {
		return nil, err
	}
	awardAmount, ok := tawardAmount.(string)
	if !ok {
		return nil, fmt.Errorf("Errors Bad type 'AwardAmount' not string")
	}
	award.AwardAmount = awardAmount

	// Organization
	tawardOrg, err := result.Find(result, []string{"Organization"}...)
	if err != nil {
		return nil, err
	}
	temp, err := json.Marshal(&tawardOrg)
	if err != nil {
		return nil, fmt.Errorf("Error in json encoding")
	}
	_, ok = tawardOrg.([]interface{})
	if !ok {
		// Unmarshal单对象
		awardOrg := &BasicOrganization{}
		err = json.Unmarshal(temp, awardOrg)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		award.Organization = append(award.Organization, awardOrg)
	} else {
		// Unmarshal列表
		err = json.Unmarshal(temp, &award.Organization)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
	}
	// Institution
	tawardIns, err := result.Find(result, []string{"Institution"}...)
	if err != nil {
		return nil, err
	}
	temp, err = json.Marshal(&tawardIns)
	if err != nil {
		return nil, fmt.Errorf("Error in json encoding")
	}
	_, ok = tawardIns.([]interface{})
	if !ok {
		// Unmarshal单对象
		awardIns := &BasicInstitutionData{}
		err = json.Unmarshal(temp, awardIns)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		award.Institution = append(award.Institution, awardIns)
	} else {
		// Unmarshal列表
		err = json.Unmarshal(temp, &award.Institution)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
	}
	// Investigator
	tawardInv, err := result.Find(result, []string{"Investigator"}...)
	if err != nil {
		return nil, err
	}
	temp, err = json.Marshal(&tawardInv)
	if err != nil {
		return nil, fmt.Errorf("Error in json encoding")
	}
	_, ok = tawardInv.([]interface{})
	if !ok {
		// Unmarshal单对象
		awardInv := &BasicInvestigatorData{}
		err = json.Unmarshal(temp, awardInv)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		award.Investigator = append(award.Investigator, awardInv)
	} else {
		// Unmarshal列表
		err = json.Unmarshal(temp, &award.Investigator)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
	}
	// reference
	tawardRef, err := result.Find(result, []string{"ProgramReference"}...)
	if err != nil {
		return nil, err
	}
	temp, err = json.Marshal(&tawardRef)
	if err != nil {
		return nil, fmt.Errorf("Error in json encoding")
	}
	_, ok = tawardRef.([]interface{})
	if !ok {
		// Unmarshal单对象
		awardRef := &BasicProgramReference{}
		err = json.Unmarshal(temp, awardRef)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		award.ProgramReference = append(award.ProgramReference, awardRef)
	} else {
		// Unmarshal列表
		err = json.Unmarshal(temp, &award.ProgramReference)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
	}
	return award, nil
}

// Part Multi Search

// AwardSearchResult Dbops层向上层传递的查询结果结构体.
type AwardSearchResult struct {
	AwardID             string        `json:"AwardID"`
	AwardTitle          string        `json:"AwardTitle"`
	AwardAmount         string        `json:"AwardAmount"`
	AwardExpirationDate string        `json:"AwardExpirationDate"`
	Last                []interface{} `json:"Last"`
}

type (
	// MultiSearchQuery 在应用内部定义的查询结构体
	MultiSearchQuery = defs.MultiSearchQuery
)

// GetMultiSearchResult 多项搜索函数->与存储层交互得到查询结果
// Params multi:查询结构体
// Params index:查询的index名称
func GetMultiSearchResult(multi *MultiSearchQuery, index string, source []string) ([]*AwardSearchResult, error) {
	var result interface{}
	award := []*AwardSearchResult{}
	ais, err := esconn.NewAwardMultiSearcher()
	if err != nil {
		return nil, err
	}
	ais.SetSource(source)
	req, err := searcher.NewMultiSearcherReq(multi)
	if err != nil {
		return nil, errors.New("Errors Bad type 'id' not type string")
	}
	tres, err := ais.Request(req, index)
	if err != nil {
		return nil, fmt.Errorf("Error Request: %s", err)
	}
	result, err = tres.Find(tres, []string{"[]", "_source", "doc", "Award"}...)
	if err != nil {
		return nil, fmt.Errorf("Error Find(): %s", err)
	}

	for _, ob := range result.([]interface{}) {
		ta := &AwardSearchResult{}
		temp, err := json.Marshal(&ob)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		err = json.Unmarshal(temp, ta)
		if err != nil {
			return nil, fmt.Errorf("Error in json decoding: %s", err)
		}
		award = append(award, ta)
	}
	olres, err := tres.Find(tres, []string{"[]", "sort"}...)
	var ok bool
	ol, ok := olres.([]interface{})
	if !ok {
		return nil, errors.New("Error in olres.([]interface{}),failed type assertion.")
	}
	for i := 0; i < len(award); i++ {
		// fmt.Printf("debug!!!:%v\n", ol[i])
		award[i].Last, ok = ol[i].([]interface{})
		if !ok {
			return nil, errors.New("Error in ol[i].([]string),failed type assertion.")
		}
	}
	return award, nil
}

type MetricVal struct {
	Value int `json:"value"`
}

type AwardAggResult struct {
	Key       string     `json:"Key"`
	DocCount  int        `json:"doc_count"`
	MetricVal *MetricVal `json:"nsf_order"`
}

type AwardAggSearchResult struct {
	Search []*AwardSearchResult
	Agg    []*AwardAggResult
	Last   []interface{}
}

func GetAggSearchResult(multi *MultiSearchQuery, index string, source []string) (*AwardAggSearchResult, error) {
	if multi.Aggs.Bucket == "" {
		return nil, fmt.Errorf("Error: Empty bucket,maybe you should try GetBucketSearchResult.")
	}
	var result searcher.SearcherResp
	award := []*AwardSearchResult{}
	agg := []*AwardAggResult{}
	ais, err := esconn.NewAwardAggSearcher()
	if err != nil {
		return nil, err
	}
	if len(source) != 0 {
		ais.SetSource(source)
	}
	req, err := searcher.NewAggSearcherReq(multi)
	if err != nil {
		return nil, errors.New("Errors Bad type 'id' not type MultiSearchQuery")
	}
	result, err = ais.Request(req, index)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("debug!!! %v\n", result)
	obs, err := result.Find(result, []string{"hits", "hits", "[]", "_source", "doc", "Award"}...)
	if err != nil {
		return nil, err
	}
	for _, ob := range obs.([]interface{}) {
		ta := &AwardSearchResult{}
		temp, err := json.Marshal(&ob)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		err = json.Unmarshal(temp, ta)
		if err != nil {
			return nil, fmt.Errorf("Error in json decoding: %s", err)
		}
		award = append(award, ta)
	}
	tar, err := result.Find(result, []string{"aggregations", "nsf_aggs", "buckets"}...)
	if err != nil {
		return nil, err
	}
	ar := tar.([]interface{})
	for _, val := range ar {
		ta := &AwardAggResult{}
		temp, err := json.Marshal(&val)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		err = json.Unmarshal(temp, ta)
		if err != nil {
			return nil, fmt.Errorf("Error in json decoding: %s", err)
		}
		agg = append(agg, ta)
	}
	la, err := result.Find(result, []string{"hits", "hits", "[]", "sort"}...)
	if err != nil {
		return nil, err
	}
	pos := len(la.([]interface{})) - 1
	last, ok := la.([]interface{})[pos].([]interface{})
	if !ok {
		return nil, errors.New("Error in last assertion.")
	}

	return &AwardAggSearchResult{
		Search: award,
		Agg:    agg,
		Last:   last,
	}, nil
}

type AwardAmountStats struct {
	Count float32 `json:"count"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Avg   float32 `json:"avg"`
	Sum   float32 `json:"sum"`
}

type AwardBucketSearchResult struct {
	Search []*AwardSearchResult
	Stats  *AwardAmountStats
}

func GetBucketSearchResult(multi *MultiSearchQuery, index string, source []string) (*AwardBucketSearchResult, error) {
	var result searcher.SearcherResp
	award := []*AwardSearchResult{}
	ais, err := esconn.NewAwardAggSearcher()
	if err != nil {
		return nil, err
	}
	if len(source) != 0 {
		ais.SetSource(source)
	}
	req, err := searcher.NewAggSearcherReq(multi)
	if err != nil {
		return nil, errors.New("Errors Bad type 'id' not type MultiSearchQuery")
	}
	result, err = ais.Request(req, index)
	if err != nil {
		return nil, err
	}
	obs, err := result.Find(result, []string{"hits", "hits", "[]", "_source", "doc", "Award"}...)
	if err != nil {
		return nil, err
	}
	for _, ob := range obs.([]interface{}) {
		ta := &AwardSearchResult{}
		temp, err := json.Marshal(&ob)
		if err != nil {
			return nil, fmt.Errorf("Error in json encoding: %s", err)
		}
		err = json.Unmarshal(temp, ta)
		if err != nil {
			return nil, fmt.Errorf("Error in json decoding: %s", err)
		}
		award = append(award, ta)
	}
	// t, _ := json.Marshal(result)
	// fmt.Printf("%s\n", t)
	aggResult, err := result.Find(result, []string{"aggregations", "nsf_aggs"}...)
	if err != nil {
		return nil, err
	}
	ta := &AwardAmountStats{}
	temp, err := json.Marshal(&aggResult)
	if err != nil {
		return nil, fmt.Errorf("Error in json encoding: %s", err)
	}
	err = json.Unmarshal(temp, ta)
	if err != nil {
		return nil, fmt.Errorf("Error in json decoding: %s", err)
	}
	return &AwardBucketSearchResult{
		Search: award,
		Stats:  ta,
	}, nil
}
