package defs

// MultiSearch structs

type AmdTimeRange struct {
	Op  int
	Gte string
	Lte string
}

type EffectTimeRange struct {
	Op  int
	Gte string
	Lte string
}

type ExpireTimeRange struct {
	Op  int
	Gte string
	Lte string
}

type AwardAmount struct {
	Op  int
	Gte int
	Lte int
}

type AwardInstrument struct {
	Op   int
	Type string
}

type AwardOrganization struct {
	Op       int
	Name     string
	Division string
}

type AwardReference struct {
	Op   int
	Name string
}

type AwardInstitution struct {
	Op      int
	Name    string
	Country string
	State   string
	City    string
}

type Sort struct {
	// 两个选项："Amount"与"Score"
	Field string
	Order string
}

type Aggs struct {
	// 两个选项：institution与reference
	Bucket string
	// 两个选项：sum与count
	// 还有"stats"是直接计算综述
	Sort string
	// true:"asc"，false:"desc"
	Order bool
}

type MultiSearchQuery struct {
	AmdRange     []*AmdTimeRange
	EffectRange  []*EffectTimeRange
	ExpireRange  []*ExpireTimeRange
	Amount       []*AwardAmount
	Instrument   []*AwardInstrument
	Organization []*AwardOrganization
	Reference    []*AwardReference
	Institution  []*AwardInstitution
	Sort         []*Sort
	Aggs         *Aggs
	All          bool
}
