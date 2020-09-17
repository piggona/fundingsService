package dbops

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/piggona/fundingsView/api/esconn"
	"github.com/piggona/fundingsView/api/esconn/searcher"
	"github.com/piggona/fundingsView/api/utils/log"
)

type FundElement struct {
	Industries          []string            `json:"industries"`
	Technology          []string            `json:"technology"`
	Description         string              `json:"description"`
	AwardEffecticeDate  string              `json:"award_effectice_date"`
	AwardExpirationDate string              `json:"award_expiration_date"`
	AwardID             string              `json:"award_id"`
	Organization        []*Organization     `json:"organization"`
	Investigator        []*Investigator     `json:"investigator"`
	ProgramReference    []*ProgramReference `json:"program_reference"`
	ProgramElement      []*ProgramElement   `json:"program_element"`
	AwardTitle          string              `json:"award_title"`
	Institution         []*Institution      `json:"institution"`
	AwardAmount         int                 `json:"award_amount"`
}

type Organization struct {
	Division    string `json:"division"`
	Code        string `json:"code"`
	Directorate string `json:"directorate"`
}

type Investigator struct {
	FirstName    string `json:"firstname"`
	Enddate      string `json:"end_date"`
	Lastname     string `json:"lastname"`
	EmailAddress string `json:"email_address"`
	StartDate    string `json:"start_date"`
	RoleCode     string `json:"role_code"`
}

type ProgramReference struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type ProgramElement struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type Institution struct {
	StateCode   string `json:"state_code"`
	StateName   string `json:"state_name"`
	Name        string `json:"name"`
	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

type SearchResultElement struct {
	Source *FundElement `json:"_source"`
}

type SearchParams struct {
	From           string
	Title          string
	TitleOp        string
	Description    string
	DescriptionOp  string
	Institution    string
	InstitutionOp  string
	Organization   string
	OrganizationOp string
	Technology     string
	TechnologyOp   string
	Industries     string
	IndustriesOp   string
	DateFrom       time.Time
	DateTo         time.Time
}

func GetMultiSearch(searchParams *SearchParams) ([]*FundElement, error) {
	bodyElement := []*SearchResultElement{}
	ais, err := esconn.NewAwardTemplateSearcher()
	if err != nil {
		log.Error("create new plain searcher error: %s", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	templateId := basicrankamount
	params := ParamsParser(searchParams)
	resp, err := ais.Request(ctx, searcher.NewTemplateSearchReq(templateId, params), INDEX)
	if err != nil {
		log.Error("searcher request error: %s", err)
		return nil, err
	}
	obj, err := resp.Find(nil, "hits", "hits")
	if err != nil {
		log.Error("resp find error: %s", err)
		return nil, err
	}
	jsonObj, ok := obj.(*simplejson.Json)
	if !ok {
		log.Error("assertion error when assert object to simplejson.Json: %s", err)
		return nil, err
	}
	objStr, err := jsonObj.Encode()
	if err != nil {
		log.Error("json object encode error: %s", err)
		return nil, err
	}
	err = json.Unmarshal(objStr, &bodyElement)
	if err != nil {
		log.Error("json string unmarshal to SearchResultElement error: %s", err)
		return nil, err
	}
	result := make([]*FundElement, len(bodyElement))
	for id, bucs := range bodyElement {
		buckets := bucs
		result[id] = buckets.Source
	}
	return result, nil
}

func ParamsParser(params *SearchParams) map[string]string {
	result := make(map[string]string)
	if _, err := strconv.Atoi(params.From); err != nil {
		result["from"] = "0"
	}
	result["title"] = params.Title
	result["description"] = params.Description
	result["institution"] = params.Institution
	result["organization"] = params.Organization
	result["technology"] = params.Technology
	result["industries"] = params.Industries
	result["date_from"] = params.DateFrom.Format("01/02/2006")
	result["date_to"] = params.DateTo.Format("01/02/2006")
	opJudge := func(op string) string {
		if op != "and" && op != "or" {
			return "and"
		}
		return op
	}
	result["title_op"] = opJudge(params.TitleOp)
	result["description_op"] = opJudge(params.DescriptionOp)
	result["institution_op"] = opJudge(params.InstitutionOp)
	result["organization_op"] = opJudge(params.OrganizationOp)
	result["technology_op"] = opJudge(params.TechnologyOp)
	result["industries_op"] = opJudge(params.IndustriesOp)
	return result
}
