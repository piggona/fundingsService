package defs

type OrganizationTitleDataResponse struct {
	Title             string `json:"title"`
	AbstractNarration string `json:"AbstractNarration"`
}

type OrganizationRankList struct {
	Rank  int    `json:"rank"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Date  int    `json:"date"`
}

type OrganizationRelatedFundResponse struct {
	RankList []*OrganizationRankList `json:"ranklist"`
}

type OrganizationRelatedTechResponse struct {
	DataOne  *GramData               `json:"dataOne"`
	DataTwo  *GramData               `json:"dataTwo"`
	RankList []*OrganizationRankList `json:"ranklist"`
}

type OrganizationRelatedOrgResponse struct {
	DataOne  *GramData               `json:"dataOne"`
	DataTwo  *GramData               `json:"dataTwo"`
	RankList []*OrganizationRankList `json:"ranklist"`
}

type OrganizationRelatedInduResponse struct {
	DataOne  *GramData               `json:"dataOne"`
	DataTwo  *GramData               `json:"dataTwo"`
	RankList []*OrganizationRankList `json:"ranklist"`
}
