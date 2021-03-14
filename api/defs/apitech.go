package defs

type TechnologyTitleDataResponse struct {
	Title             string `json:"title"`
	AbstractNarration string `json:"AbstractNarration"`
}

type TechnologyRankList struct {
	Rank  int    `json:"rank"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Date  int    `json:"date"`
}

type TechnologyRelatedFundResponse struct {
	RankList []*TechnologyRankList `json:"ranklist"`
}

type TechnologyRelatedDivResponse struct {
	DataOne  *GramData             `json:"dataOne"`
	DataTwo  *GramData             `json:"dataTwo"`
	RankList []*TechnologyRankList `json:"ranklist"`
}

type TechnologyRelatedOrgResponse struct {
	DataOne  *GramData             `json:"dataOne"`
	DataTwo  *GramData             `json:"dataTwo"`
	RankList []*TechnologyRankList `json:"ranklist"`
}

type TechnologyRelatedInduResponse struct {
	DataOne  *GramData             `json:"dataOne"`
	DataTwo  *GramData             `json:"dataTwo"`
	RankList []*TechnologyRankList `json:"ranklist"`
}
