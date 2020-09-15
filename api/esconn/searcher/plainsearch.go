package searcher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/piggona/fundingsView/api/esconn/connector"
	"github.com/piggona/fundingsView/api/utils/log"
)

type PlainSearch struct {
	conn   connector.Conn
	source []string
}

type PlainSearchReq string

func (p *PlainSearchReq) Validate() (interface{}, error) {
	json, err := simplejson.NewJson([]byte(*p))
	if err != nil {
		log.Error("json parse error: %s", err)
	}
	return json, nil
}

func NewPlainSearchReq(req *string) *PlainSearchReq {
	if req == nil {
		log.Error("empty request body error")
		return nil
	}
	plain := PlainSearchReq(*req)
	return &plain
}

type PlainSearchResp simplejson.Json

func (p *PlainSearchResp) Find(resp interface{}, paths ...string) (interface{}, error) {
	if resp == nil {
		json := simplejson.Json(*p)
		return json.GetPath(paths...), nil
	}
	if json, ok := resp.(simplejson.Json); ok {
		return json.GetPath(paths...), nil
	}
	if len(paths) == 0 {
		return resp, nil
	}
	log.Error("resp type not correct")
	return nil, fmt.Errorf("resp type not correct")
}

func NewPlainSearchResp(resp interface{}) (*PlainSearchResp, error) {
	json, err := simplejson.NewJson([]byte(resp.(string)))
	if err != nil {
		log.Error("simplejson initialization failed: %s", err)
		return nil, err
	}
	ps := PlainSearchResp(*json)
	return &ps, nil
}

func (p *PlainSearch) SetConn(conn connector.Conn) {
	p.conn = conn
}

func (p *PlainSearch) SetSource(source []string) {
	p.source = source
}

func (p *PlainSearch) Request(ctx context.Context, input SearcherReq, index string) (SearcherResp, error) {
	searchmap, err := p.parser(input)
	if err != nil {
		log.Error("parser error: %s", err)
		return nil, err
	}
	res, err := p.conn.Query(ctx, searchmap, index)
	if err != nil {
		log.Error("query error: %s", err)
		return nil, err
	}
	resStr, _ := json.Marshal(res)
	result, err := NewPlainSearchResp(resStr)
	if err != nil {
		log.Error("plain response creation error: %s", err)
		return nil, err
	}
	return result, nil
}

func (p *PlainSearch) parser(input SearcherReq) (map[string]interface{}, error) {
	i, ok := input.(*PlainSearchReq)
	if !ok {
		log.Error("type assertion failed")
		return nil, fmt.Errorf("Error bad request type, expect type: *PlainSearchReq")
	}
	obj, err := i.Validate()
	if err != nil {
		log.Error("validation error: %s", err)
		return nil, err
	}
	json := obj.(simplejson.Json)
	inputMap, err := json.Map()
	if err != nil {
		log.Error("json convert to map failed: %s", err)
		return nil, fmt.Errorf("json convert to map failed: %s", err)
	}
	if len(p.source) != 0 {
		inputMap["_source"] = p.source
	}

	return inputMap, nil
}
