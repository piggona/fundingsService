package searcher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/piggona/fundingsView/api/esconn/connector"
	"github.com/piggona/fundingsView/api/utils/log"
)

type TemplateSearch struct {
	conn   connector.Conn
	source []string
}

type TemplateSearchReq struct {
	Id     string            `json:"id"`
	Params map[string]string `json:"params"`
}

func (p *TemplateSearchReq) Validate() (interface{}, error) {
	request := make(map[string]interface{})
	request["id"] = p.Id
	request["params"] = p.Params
	return request, nil
}

func NewTemplateSearchReq(templateId string, params map[string]string) *TemplateSearchReq {
	return &TemplateSearchReq{
		Id:     templateId,
		Params: params,
	}
}

type TemplateSearchResp simplejson.Json

func (p *TemplateSearchResp) Find(resp interface{}, paths ...string) (interface{}, error) {
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

func NewTemplateSearchResp(resp interface{}) (*TemplateSearchResp, error) {
	json, err := simplejson.NewJson(resp.([]byte))
	if err != nil {
		log.Error("simplejson initialization failed: %s", err)
		return nil, err
	}
	ps := TemplateSearchResp(*json)
	return &ps, nil
}

func (p *TemplateSearch) SetConn(conn connector.Conn) {
	p.conn = conn
}

func (p *TemplateSearch) SetSource(source []string) {
	p.source = source
}

func (p *TemplateSearch) Request(ctx context.Context, input SearcherReq, index string) (SearcherResp, error) {
	searchmap, err := p.parser(input)
	if err != nil {
		log.Error("parser error: %s", err)
		return nil, err
	}

	res, err := p.conn.QueryTemplate(ctx, searchmap, index)
	if err != nil {
		log.Error("query error: %s", err)
		return nil, err
	}
	resStr, _ := json.Marshal(res)
	result, err := NewTemplateSearchResp(resStr)
	if err != nil {
		log.Error("template response creation error: %s", err)
		return nil, err
	}
	return result, nil
}

func (p *TemplateSearch) parser(input SearcherReq) (map[string]interface{}, error) {
	i, ok := input.(*TemplateSearchReq)
	if !ok {
		log.Error("type assertion failed")
		return nil, fmt.Errorf("Error bad request type, expect type: *TemplateSearchReq")
	}
	obj, err := i.Validate()
	if err != nil {
		log.Error("validation error: %s", err)
		return nil, err
	}

	return obj.(map[string]interface{}), nil
}
