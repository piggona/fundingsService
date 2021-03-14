package defs

type IndustryTitleDataResponse struct {
	Title             string `json:"title"`
	AbstractNarration string `json:"AbstractNarration"`
}

type IndustryRankList struct {
	Rank  int    `json:"rank"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Date  int    `json:"date"`
}

type IndustryRelatedFundResponse struct {
	RankList []*IndustryRankList `json:"ranklist"`
}

type IndustryRelatedTechResponse struct {
	DataOne  *GramData           `json:"dataOne"`
	DataTwo  *GramData           `json:"dataTwo"`
	RankList []*IndustryRankList `json:"ranklist"`
}

type IndustryRelatedOrgResponse struct {
	DataOne  *GramData           `json:"dataOne"`
	DataTwo  *GramData           `json:"dataTwo"`
	RankList []*IndustryRankList `json:"ranklist"`
}

type IndustryRelatedInduResponse struct {
	DataOne  *GramData           `json:"dataOne"`
	DataTwo  *GramData           `json:"dataTwo"`
	RankList []*IndustryRankList `json:"ranklist"`
}
