package connector

// Conn 与elasticsearch实例的连接接口.
type Conn interface {
	// Query 接收用户请求（map），使用句柄向elasticsearch实例发起请求，并将responseBody解析为map返回
	Query(req map[string]interface{}, index string) (map[string]interface{}, error)
	// getHandler() 获取可用句柄.句柄是类定义的连接elasticsearch实例的方式，返回的可以是es的操作对象
	// 也可以是channel等代表与elasticsearch通信的操作对象
	getHandler() (interface{}, error)
}
