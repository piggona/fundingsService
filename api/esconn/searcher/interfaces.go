package searcher

import "github.com/piggona/fundingsView/api/esconn/connector"

// Searcher 用户定义的es查询对象.
type Searcher interface {
	// Request 接收用户请求，使用parser将用户请求转化为DSL（map形式）
	// 之后使用conn与elasticsearch进行交互,得到结果绑定到指定的output结构体中
	Request(input SearcherReq, index string) (SearcherResp, error)
	// 自定义对象的Conn
	SetConn(conn connector.Conn)
	SetSource(source []string)
	parser(input SearcherReq) (map[string]interface{}, error)
}

type SearcherReq interface {
	Validate() (interface{}, error)
}

type SearcherResp interface {
	Find(resp interface{}, paths ...string) (interface{}, error)
}
