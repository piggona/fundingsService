package defs

type CategoryTitleDataResponse struct {
	Title             string `json:"title"`
	AbstractNarration string `json:"AbstractNarration"`
}

type CategoryRankList struct {
	Rank  int    `json:"rank"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Date  int    `json:"date"`
}

type CategoryRelatedFundResponse struct {
	RankList []*CategoryRankList `json:"ranklist"`
}

type GramDataBucket struct {
	Name string   `json:"name"`
	Data []string `json:"data"`
}

type GramData struct {
	XAxis []string          `json:"xAxis"`
	Data  []*GramDataBucket `json:"data"`
}

type CategoryRelatedTechResponse struct {
	DataOne  *GramData           `json:"dataOne"`
	DataTwo  *GramData           `json:"dataTwo"`
	RankList []*CategoryRankList `json:"ranklist"`
}

type CategoryRelatedOrgResponse struct {
	DataOne  *GramData           `json:"dataOne"`
	DataTwo  *GramData           `json:"dataTwo"`
	RankList []*CategoryRankList `json:"ranklist"`
}

type CategoryRelatedInduResponse struct {
	DataOne  *GramData           `json:"dataOne"`
	DataTwo  *GramData           `json:"dataTwo"`
	RankList []*CategoryRankList `json:"ranklist"`
}
