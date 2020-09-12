package searcher

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/piggona/fundingsView/api/esconn/connector"
)

// AwardIDSearch 搜索执行结构体
type AwardIDSearch struct {
	conn   connector.Conn
	source []string
}

func (a *AwardIDSearch) parser(input SearcherReq) (map[string]interface{}, error) {
	i, ok := input.(*IDSearcherReq)
	if !ok {
		return nil, fmt.Errorf("Error bad request type, expect type: *IDSearcherReq")
	}
	str, err := i.Validate()
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"_source": a.source,
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"doc.Award.AwardID.keyword": str.(string),
			},
		},
	}
	return result, nil
}

// SetSource 设置返回的字段类型
func (a *AwardIDSearch) SetSource(source []string) {
	a.source = source
}

// SetConn 设置与es的连接
func (a *AwardIDSearch) SetConn(conn connector.Conn) {
	a.conn = conn
}

// Request 将用户传入的数据转换数据形式，执行elasticsearch搜索，将结果数据转换到指定结构体
func (a *AwardIDSearch) Request(input SearcherReq, index string) (SearcherResp, error) {
	searchmap, err := a.parser(input)
	if err != nil {
		return nil, err
	}
	res, err := a.conn.Query(searchmap, index)
	if err != nil {
		return nil, err
	}
	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["doc"].(map[string]interface{})["Award"]

	result, err := NewIDSearcherResp(source)
	if err != nil {
		return nil, err
	}
	// output := &AwardIDResult{
	// 	AwardID:             source.(map[string]interface{})["AwardID"].(string),
	// 	AwardTitle:          source.(map[string]interface{})["AwardTitle"].(string),
	// 	AwardAmount:         source.(map[string]interface{})["AwardAmount"].(string),
	// 	AwardExpirationDate: source.(map[string]interface{})["AwardExpirationDate"].(string),
	// }
	return result, nil
}

type IDSearcherReq string

func (i *IDSearcherReq) Validate() (interface{}, error) {
	return string(*i), nil
}

func NewIDSearcherReq(req string) (*IDSearcherReq, error) {
	id := IDSearcherReq(req)
	return &id, nil
}

type IDSearcherResp map[string]interface{}

func (idsearch *IDSearcherResp) Find(resp interface{}, paths ...string) (interface{}, error) {
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
			tt, ok := temp.(*IDSearcherResp)
			if !ok {
				tm, ok := temp.(map[string]interface{})
				if !ok {
					return nil, errors.New("Error in type assertion map[string]interface{}")
				}
				temp = tm[path]
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

func NewIDSearcherResp(resp interface{}) (*IDSearcherResp, error) {
	res, ok := resp.(map[string]interface{})
	if !ok {
		return nil, errors.New("Error bad type resp.")
	}
	result := IDSearcherResp(res)
	return &result, nil
}
