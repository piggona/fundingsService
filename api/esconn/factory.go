package esconn

import (
	"github.com/piggona/fundingsView/api/esconn/connector"
	"github.com/piggona/fundingsView/api/esconn/searcher"
)

// NewAwardIDSearcher 基金ID精确搜索执行对象工厂函数
func NewAwardIDSearcher() (searcher.Searcher, error) {
	source := []string{"doc.Award.AwardID", "doc.Award.AwardTitle", "doc.Award.AwardAmount", "doc.Award.AwardExpirationDate"}
	var err error
	conn, err := connector.NewOrdinalConn()
	if err != nil {
		return nil, err
	}
	result := &searcher.AwardIDSearch{}
	result.SetConn(conn)
	result.SetSource(source)
	return result, nil
}

// NewAwardMultiSearcher 基金多项搜索执行对象工厂函数
func NewAwardMultiSearcher() (searcher.Searcher, error) {
	source := []string{"doc.Award.AwardID", "doc.Award.AwardTitle", "doc.Award.AwardAmount", "doc.Award.AwardExpirationDate"}
	var err error
	conn, err := connector.NewOrdinalConn()
	if err != nil {
		return nil, err
	}
	result := &searcher.AwardMultiSearch{}
	result.SetConn(conn)
	result.SetSource(source)
	return result, nil
}

func NewAwardAggSearcher() (searcher.Searcher, error) {
	source := []string{"doc.Award.AwardID", "doc.Award.AwardTitle", "doc.Award.AwardAmount", "doc.Award.AwardExpirationDate"}
	var err error
	conn, err := connector.NewOrdinalConn()
	if err != nil {
		return nil, err
	}
	result := &searcher.AwardAggSearch{}
	result.SetConn(conn)
	result.SetSource(source)
	return result, nil
}
