package searcher

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/piggona/fundingsView/api/defs"
	"github.com/piggona/fundingsView/api/esconn/connector"
)

type (
	MultiSearchQuery  = defs.MultiSearchQuery
	AmdTimeRange      = defs.AmdTimeRange
	EffectTimeRange   = defs.EffectTimeRange
	ExpireTimeRange   = defs.ExpireTimeRange
	AwardAmount       = defs.AwardAmount
	AwardInstrument   = defs.AwardInstrument
	AwardOrganization = defs.AwardOrganization
	AwardReference    = defs.AwardReference
	AwardInstitution  = defs.AwardInstitution
)

// AwardMultiSearch 搜索执行结构体
type AwardMultiSearch struct {
	conn   connector.Conn
	source []string
}

// parser 将用户输入经过 判断类型有效（assert）-> 管道处理（validate）-> 编码后传递给Request发出
func (a *AwardMultiSearch) parser(input SearcherReq) (map[string]interface{}, error) {
	// 先判断用户的输入是否符合本类标准输入（MultiSearchReq），然后使用Validate对输入做处理并转为普遍类型。
	_, ok := input.(*MultiSearcherReq)
	if !ok {
		return nil, fmt.Errorf("Error bad request type, expect type: *MultiSearcherReq")
	}
	tti, err := input.Validate()
	if err != nil {
		return nil, err
	}
	i := tti.(defs.MultiSearchQuery)

	var result map[string]interface{}
	// 配置搜索
	if i.All {
		result = map[string]interface{}{
			"_source": a.source,
			"size":    20,
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
			"aggs": map[string]interface{}{},
			"sort": []interface{}{},
		}
	} else {
		result = map[string]interface{}{
			"_source": a.source,
			"size":    20,
			"query": map[string]interface{}{
				"bool": map[string][]interface{}{
					"must":       {},
					"must_not":   {},
					"should":     {},
					"should_not": {},
				},
			},
			"aggs": map[string]interface{}{},
			"sort": []interface{}{},
		}
		put_in_right_place := func(op int, str map[string]interface{}) {
			switch op {
			case 0:
				t := result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["must"]
				result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["must"] = append(t, str)
			case 1:
				t := result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["must_not"]
				result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["must_not"] = append(t, str)
			case 2:
				t := result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["should"]
				result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["should"] = append(t, str)
			case 3:
				t := result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["should_not"]
				result["query"].(map[string]interface{})["bool"].(map[string][]interface{})["should_not"] = append(t, str)
			}
		}
		amd := i.AmdRange
		for _, val := range amd {
			temp := map[string]interface{}{
				"range": map[string]interface{}{
					"doc.Award.MinAmdLetterDate": map[string]interface{}{
						"gte": val.Gte,
						"lte": val.Lte,
					},
				},
			}
			put_in_right_place(val.Op, temp)
		}
		eft := i.EffectRange
		for _, val := range eft {
			temp := map[string]interface{}{
				"range": map[string]interface{}{
					"doc.Award.AwardEffectiveDate": map[string]interface{}{
						"gte": val.Gte,
						"lte": val.Lte,
					},
				},
			}
			put_in_right_place(val.Op, temp)
		}
		exp := i.ExpireRange
		for _, val := range exp {
			temp := map[string]interface{}{
				"range": map[string]interface{}{
					"doc.Award.AwardExpirationDate": map[string]interface{}{
						"gte": val.Gte,
						"lte": val.Lte,
					},
				},
			}
			put_in_right_place(val.Op, temp)
		}
		amount := i.Amount
		for _, val := range amount {
			temp := map[string]interface{}{
				"range": map[string]interface{}{
					"doc.Award.AwardAmount": map[string]interface{}{
						"gte": val.Gte,
						"lte": val.Lte,
					},
				},
			}
			put_in_right_place(val.Op, temp)
		}
		tru := i.Instrument
		for _, val := range tru {
			temp := map[string]interface{}{
				"term": map[string]interface{}{
					"doc.Award.AwardInstrument.Value.keyword": val.Type,
				},
			}
			put_in_right_place(val.Op, temp)
		}
		org := i.Organization
		for _, val := range org {
			if len(val.Name) != 0 {
				temp := map[string]interface{}{
					"match": map[string]interface{}{
						"doc.Award.Organization.Directorate.LongName": val.Name,
					},
				}
				put_in_right_place(val.Op, temp)
			}
			if len(val.Division) != 0 {
				temp := map[string]interface{}{
					"match": map[string]interface{}{
						"doc.Award.Organization.Division.LongName": val.Division,
					},
				}
				put_in_right_place(val.Op, temp)
			}
		}
		ins := i.Institution
		for _, val := range ins {
			if len(val.Name) != 0 {
				temp := map[string]interface{}{
					"match": map[string]interface{}{
						"doc.Award.Institution.Name": val.Name,
					},
				}
				put_in_right_place(val.Op, temp)
			}
			if len(val.Country) != 0 {
				temp := map[string]interface{}{
					"match": map[string]interface{}{
						"doc.Award.Institution.CountryName": val.Country,
					},
				}
				put_in_right_place(val.Op, temp)
			}
			if len(val.State) != 0 {
				temp := map[string]interface{}{
					"match": map[string]interface{}{
						"doc.Award.Institution.StateName": val.State,
					},
				}
				put_in_right_place(val.Op, temp)
			}
			if len(val.City) != 0 {
				temp := map[string]interface{}{
					"match": map[string]interface{}{
						"doc.Award.Institution.CityName": val.City,
					},
				}
				put_in_right_place(val.Op, temp)
			}
		}
	}
	// 配置sort
	for _, val := range i.Sort {
		var temp map[string]interface{}
		switch val.Field {
		case "Amount":
			temp = map[string]interface{}{
				"doc.Award.AwardAmount": map[string]interface{}{
					"order": val.Order,
				},
			}
		case "Score":
			temp = map[string]interface{}{
				"_score": map[string]interface{}{
					"order": val.Order,
				},
			}
		default:
			continue
		}
		result["sort"] = append(result["sort"].([]interface{}), temp)
	}
	// r, _ := json.Marshal(result)
	// fmt.Printf("%s", r)
	return result, nil
}

// SetSource 设置返回的字段类型
func (a *AwardMultiSearch) SetSource(source []string) {
	a.source = source
}

// SetConn 设置与es的连接
func (a *AwardMultiSearch) SetConn(conn connector.Conn) {
	a.conn = conn
}

// Request 将用户传入的数据转换数据形式，执行elasticsearch搜索，将结果数据转换到指定结构体
func (a *AwardMultiSearch) Request(input SearcherReq, index string) (SearcherResp, error) {
	searchmap, err := a.parser(input)
	if err != nil {
		return nil, err
	}
	res, err := a.conn.Query(searchmap, index)
	if err != nil {
		return nil, err
	}
	source := res["hits"].(map[string]interface{})["hits"]
	result, err := NewMultiSearcherResp(source)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MultiSearcherReq 用户标准输入，满足SearcherReq接口
type MultiSearcherReq defs.MultiSearchQuery

// Validate 管道处理，将标准输入转为对外可用类型
func (i *MultiSearcherReq) Validate() (interface{}, error) {
	return defs.MultiSearchQuery(*i), nil
}

// NewMultiSearcherReq 从一个defs.MultiSearchQuery对象，创建标准输入对象
func NewMultiSearcherReq(req *MultiSearchQuery) (*MultiSearcherReq, error) {
	id := MultiSearcherReq(*req)
	return &id, nil
}

// MultiSearcherResp 标准输出，满足SearcherResp接口
type MultiSearcherResp []interface{}

// Find 标准输出的取数据方法，用于获取指定路径上的数据
// Param:resp 从数据的哪里开始寻找，也可以说从哪个数据对象（子对象）开始寻找
// Param:paths 数据路径
func (idsearch *MultiSearcherResp) Find(resp interface{}, paths ...string) (interface{}, error) {
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
				// fmt.Printf("debug!!! %v\n",temp)
				tt, ok := temp.(*MultiSearcherResp)
				if !ok {
					tti, ok := temp.([]interface{})
					if !ok {
						return nil, errors.New("Error in type assertion []interface{}")
					}
					for _, val := range tti {
						item, err := idsearch.Find(val, paths[i+1:]...)
						if err != nil {
							return nil, err
						}
						tres = append(tres, item)
					}
					return tres, nil
				}
				for _, val := range *tt {
					item, err := idsearch.Find(val, paths[i+1:]...)
					if err != nil {
						return nil, err
					}
					tres = append(tres, item)
				}
				return tres, nil
			}
			tt, ok := temp.(map[string]interface{})
			if !ok {
				return nil, errors.New("Error in type assertion map[string]interface{}")
			}
			temp = tt[path]
		} else {
			// 数字
			tt, ok := temp.(*MultiSearcherResp)
			if !ok {
				tti, ok := temp.([]interface{})
				if !ok {
					return nil, errors.New("Error in type assertion []interface{}")
				}
				temp = tti[t]
			}
			temp = (*tt)[t]
		}
	}
	return temp, nil
}

// NewMultiSearcherResp 从一个[]interface{},创建一个标准输出
func NewMultiSearcherResp(resp interface{}) (*MultiSearcherResp, error) {
	res, ok := resp.([]interface{})
	if !ok {
		return nil, errors.New("Error bad type resp.")
	}
	result := MultiSearcherResp(res)
	return &result, nil
}
