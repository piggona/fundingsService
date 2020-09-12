package defs

// Request Structs Definition

type UuidRequest struct {
	Uuid string `form:"uuid" json:"uuid" binding:"required"`
}

type SimilarRequest struct {
	Uuid string `form:"uuid" json:"uuid" binding:"required"`
	Page int    `form:"page" json:"page" binding:"required"`
}

// Response Structs Definition

// Response detail part
type BasicOrganizationData struct {
	LongName     string `json:"LongName"`
	Abbreviation string `json:"Abbreviation"`
}

type BasicOrganization struct {
	Directorate BasicOrganizationData `json:"Directorate"`
	Division    BasicOrganizationData `json:"Division"`
	Code        string                `json:"Code"`
}

type BasicInstitutionData struct {
	Name string `json:"Name"`
}

type BasicInvestigatorData struct {
	FullName  string `json:"FullName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type BasicProgramReference struct {
	Code string `json:"Code"`
	Text string `json:"Text"`
}

type BasicDetailResponse struct {
	AwardTitle        string                   `json:"AwardTitle"`
	AbstractNarration string                   `json:"AbstractNarration"`
	Organization      []*BasicOrganization     `json:"Organization"`
	AwardAmount       string                   `json:"AwardAmount"`
	Institution       []*BasicInstitutionData  `json:"Institution"`
	Investigator      []*BasicInvestigatorData `json:"Investigator"`
	ProgramReference  []*BasicProgramReference `json:"ProgramElement"`
	Description       string                   `json:"description"`
}

// Response CopTree WordTree part
type LevelData struct {
	Name     string       `json:"name"`
	Value    int          `json:"value"`
	Color    string       `json:"color"`
	Children []*LevelData `json:"children"`
}

type LevelDataResponse struct {
	Data []*LevelData `json:"data"`
}
