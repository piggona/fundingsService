package defs

// Request Structs Definition

// Search request
type Query struct {
	Attr  string `form:"attr" json:"attr" binding:"required"`
	Query string `form:"query" json:"query" binding:"required"`
	Logic string `form:"logic" json:"logic" binding:"required"`
}
type SearchRequest struct {
	Type  string   `form:"type" json:"type" binding:"required"`
	Sort  string   `form:"sort" json:"sort" binding:"required"`
	From  int      `form:"from" json:"from" binding:"required"`
	Size  int      `form:"size" json:"size" binding:"required"`
	Array []*Query `form:"query" json:"query" binding:"required"`
}

// Response Struct Definition

// Response Statistic Part
type Statistic struct {
	FundingsAmount float32 `json:"fundingsAmount"`
	AvgInvested    float32 `json:"avgInvested"`
}

// Response FundRank part & FundRank response
type FundRankData struct {
	Rank  string `json:"rank"`
	Name  string `json:"name"`
	Money string `json:"_money"`
}

type FundRankResponse struct {
	Data      []*FundRankData `json:"data"`
	OrderType string          `json:"order_type"`
	Last      []string        `json:"last"`
}

// Response ResPie part
type SeriesData struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type ResPieResponse struct {
	Legend []string      `json:"legend"`
	Series []*SeriesData `json:"series"`
}

// Response resCloud part

type ResCloudResponse struct {
	Data []*SeriesData `json:"data"`
}

// Response TreeNode part
type TreeTitle struct {
	Content string `json:"content"`
	Uuid    string `json:"uuid"`
}

type TreeData struct {
	Title   TreeTitle        `json:"title"`
	Depth   int              `json:"depth"`
	Phrase  string           `json:"phrase"`
	Content string           `json:"content"`
	Color   string           `json:"color"`
	Next    TreeNodeResponse `json:"next"`
}

type TreeNodeResponse struct {
	Data []*TreeData `json:"data"`
}

// Response basic part
type BasicResponse struct {
	Statistic Statistic        `json:"statistic"`
	FundRank  FundRankResponse `json:"fundRank"`
	ResPie    ResPieResponse   `json:"resPie"`
	ResCloud  ResCloudResponse `json:"resCloud"`
	TreeNodes TreeNodeResponse `json:"treeNodes"`
}

// Response Search part
// type SearchResp map[string]interface{}

// func (resp *SearchResp) GetSearchResult(source []string) (interface{}, error) {
// 	for _, s := range source {
// 		data, ok := (*resp)[s]
// 		if !ok {
// 			err := fmt.Errorf("Error no field: %s", s)
// 			return nil, err
// 		}

// 	}
// }
