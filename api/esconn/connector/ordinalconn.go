package connector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

var (
	once sync.Once
	// Ordinal 单例的连接句柄.
	Ordinal *OrdinalConn
)

// OrdinalConn 简单的elasticsearch handler
type OrdinalConn struct {
	conn *elasticsearch.Client
	info map[string]interface{}
}

// Query 使用Conn来处理请求，并将response Body解析为map返回.
// Params req:elasticsearch DSL request
// Params index:elasticsearch index used for this transaction
func (o *OrdinalConn) Query(req map[string]interface{}, index string) (map[string]interface{}, error) {
	c, err := o.getHandler()
	result := make(map[string]interface{})
	var buf bytes.Buffer
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	es, ok := c.(*elasticsearch.Client)
	if !ok {
		return nil, fmt.Errorf("Error: handler not in type *elasticsearch.Client")
	}
	if err := json.NewEncoder(&buf).Encode(&req); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing response body in res.IsError(): %s", err)
		}
		return nil, fmt.Errorf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"])
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error parsing response body: %s", err)
	}
	log.Printf("[%s] %d hits; took: %dms",
		res.Status(),
		int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(result["took"].(float64)))

	return result, nil
}

func (o *OrdinalConn) getHandler() (interface{}, error) {
	if o.conn == nil {
		return nil, fmt.Errorf("no handlers there in OrdinalConn.conn!\n")
	}
	return o.conn, nil
}

// NewOrdinalConn 单例创建connection
func NewOrdinalConn() (Conn, error) {
	info := make(map[string]interface{})
	// var es *elasticsearch.Client
	var err error
	once.Do(func() {
		var err error
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://10.103.240.34:8070/es/",
				// "http://es03:9200",
			},
		}
		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			err = fmt.Errorf("Error Creating Client: %s", err)
		}

		res, err := es.Info()
		if err != nil {
			err = fmt.Errorf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		if res.IsError() {
			err = fmt.Errorf("Error: %s", res.String())
		}
		if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
			err = fmt.Errorf("Error parsing the response body: %s", err)
		}
		// Print client and server version numbers.
		log.Printf("Client: %s", elasticsearch.Version)
		log.Printf("Server: %s", info["version"].(map[string]interface{})["number"])
		log.Println(strings.Repeat("~", 37))
		Ordinal = &OrdinalConn{
			conn: es,
			info: info,
		}
	})
	if err != nil {
		return nil, err
	}
	return Ordinal, nil
}
