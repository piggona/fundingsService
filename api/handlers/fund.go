package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/piggona/fundingsView/api/defs"
	"github.com/piggona/fundingsView/api/utils/log"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/dbops"
)

func GetDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := dbops.GetFundDetail(uuid)
	if err != nil {
		fmt.Errorf("Error get detail: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func GetCopTree(c *gin.Context) {
	uuid := c.Param("uuid")
	fund, err := dbops.GetFundDetail(uuid)
	if err != nil {
		log.Error("dbops GetFundDetail error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	orgs := make([]*defs.LevelData, len(fund.Institution))
	for id, f := range fund.Institution {
		indu := &OrgTreeVal{
			Key:   f.Name,
			Value: 50,
		}
		orgs[id] = getOrgTree(indu, 3)
	}
	resp := &defs.LevelDataResponse{
		Data: orgs,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

func GetWordTree(c *gin.Context) {
	uuid := c.Param("uuid")
	fund, err := dbops.GetFundDetail(uuid)
	if err != nil {
		log.Error("dbops GetFundDetail error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	indus := make([]*defs.LevelData, len(fund.Industries))
	for id, f := range fund.Industries {
		indu := &InduTreeVal{
			Key:   f,
			Value: 50,
		}
		indus[id] = getInduTree(indu, 3)
	}
	resp := &defs.LevelDataResponse{
		Data: indus,
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp:   resp,
	})
	return
}

type InduTreeVal struct {
	Key   string
	Value int
}

func getInduTree(induVal *InduTreeVal, depth int) *defs.LevelData {
	data := &defs.LevelData{
		Name:  induVal.Key,
		Value: induVal.Value,
		Type:  "industry",
		UUID:  induVal.Key,
	}
	depth--
	if depth <= 0 {
		return data
	}
	data.Children = []*defs.LevelData{}
	org, err := dbops.GetFundInduOrg(induVal.Key)
	if err != nil {
		log.Error("dbops GetFundInduOrg error: %s", err)
	} else {
		for _, o := range org {
			orgVal := &OrgTreeVal{
				Key:   o.Key,
				Value: o.Value,
			}
			data.Children = append(data.Children, getOrgTree(orgVal, depth))
		}
	}
	tech, err := dbops.GetFundInduTech(induVal.Key)
	if err != nil {
		log.Error("dbops GetFundInduTech error: %s", err)
	} else {
		for _, t := range tech {
			techVal := &TechTreeVal{
				Key:   t.Key,
				Value: t.Value,
			}
			data.Children = append(data.Children, getTechTree(techVal, depth))
		}
	}
	div, err := dbops.GetFundInduDiv(induVal.Key)
	if err != nil {
		log.Error("dbops GetFundInduDiv error: %s", err)
	} else {
		for _, d := range div {
			divVal := &DivTreeVal{
				Key:   d.Key,
				Name:  d.Name,
				Value: d.Value,
			}
			data.Children = append(data.Children, getDivTree(divVal, depth))
		}
	}
	return data
}

type TechTreeVal struct {
	Key   string
	Value int
}

func getTechTree(techVal *TechTreeVal, depth int) *defs.LevelData {
	data := &defs.LevelData{
		Name:  techVal.Key,
		Type:  "tech",
		UUID:  techVal.Key,
		Value: techVal.Value,
	}
	depth--
	if depth <= 0 {
		return data
	}
	data.Children = []*defs.LevelData{}
	org, err := dbops.GetFundTechOrg(techVal.Key)
	if err != nil {
		log.Error("dbops GetFundTechOrg error: %s", err)
	} else {
		for _, o := range org {
			orgVal := &OrgTreeVal{
				Key:   o.Key,
				Value: o.Value,
			}
			data.Children = append(data.Children, getOrgTree(orgVal, depth))
		}
	}
	indu, err := dbops.GetFundTechIndu(techVal.Key)
	if err != nil {
		log.Error("dbops GetFundTechTech error: %s", err)
	} else {
		for _, i := range indu {
			induVal := &InduTreeVal{
				Key:   i.Key,
				Value: i.Value,
			}
			data.Children = append(data.Children, getInduTree(induVal, depth))
		}
	}
	div, err := dbops.GetFundTechDiv(techVal.Key)
	if err != nil {
		log.Error("dbops GetFundTechDiv error: %s", err)
	} else {
		for _, d := range div {
			divVal := &DivTreeVal{
				Key:   d.Key,
				Name:  d.Name,
				Value: d.Value,
			}
			data.Children = append(data.Children, getDivTree(divVal, depth))
		}
	}
	return data
}

type DivTreeVal struct {
	Key   string
	Name  string
	Value int
}

func getDivTree(divVal *DivTreeVal, depth int) *defs.LevelData {
	data := &defs.LevelData{
		Name:  divVal.Name,
		UUID:  divVal.Key,
		Type:  "category",
		Value: divVal.Value,
	}
	depth--
	if depth <= 0 {
		return data
	}
	data.Children = []*defs.LevelData{}
	org, err := dbops.GetFundDivOrg(divVal.Key)
	if err != nil {
		log.Error("dbops GetFundDivOrg error: %s", err)
	} else {
		for _, o := range org {
			orgVal := &OrgTreeVal{
				Key:   o.Key,
				Value: o.Value,
			}
			data.Children = append(data.Children, getOrgTree(orgVal, depth))
		}
	}
	indu, err := dbops.GetFundDivIndu(divVal.Key)
	if err != nil {
		log.Error("dbops GetFundDivTech error: %s", err)
	} else {
		for _, i := range indu {
			induVal := &InduTreeVal{
				Key:   i.Key,
				Value: i.Value,
			}
			data.Children = append(data.Children, getInduTree(induVal, depth))
		}
	}
	tech, err := dbops.GetFundDivTech(divVal.Key)
	if err != nil {
		log.Error("dbops GetFundDivTech error: %s", err)
	} else {
		for _, t := range tech {
			techVal := &TechTreeVal{
				Key:   t.Key,
				Value: t.Value,
			}
			data.Children = append(data.Children, getTechTree(techVal, depth))
		}
	}
	return data
}

type OrgTreeVal struct {
	Key   string
	Value int
}

func getOrgTree(orgVal *OrgTreeVal, depth int) *defs.LevelData {
	data := &defs.LevelData{
		Name:  orgVal.Key,
		UUID:  orgVal.Key,
		Type:  "org",
		Value: orgVal.Value,
	}
	depth--
	if depth <= 0 {
		return data
	}
	data.Children = []*defs.LevelData{}
	div, err := dbops.GetFundOrgDiv(orgVal.Key)
	if err != nil {
		log.Error("dbops GetFundOrgDiv error: %s", err)
	} else {
		for _, d := range div {
			divVal := &DivTreeVal{
				Key:   d.Key,
				Name:  d.Name,
				Value: d.Value,
			}
			data.Children = append(data.Children, getDivTree(divVal, depth))
		}
	}
	indu, err := dbops.GetFundOrgIndu(orgVal.Key)
	if err != nil {
		log.Error("dbops GetFundOrgTech error: %s", err)
	} else {
		for _, i := range indu {
			induVal := &InduTreeVal{
				Key:   i.Key,
				Value: i.Value,
			}
			data.Children = append(data.Children, getInduTree(induVal, depth))
		}
	}
	tech, err := dbops.GetFundOrgTech(orgVal.Key)
	if err != nil {
		log.Error("dbops GetFundOrgTech error: %s", err)
	} else {
		for _, t := range tech {
			techVal := &TechTreeVal{
				Key:   t.Key,
				Value: t.Value,
			}
			data.Children = append(data.Children, getTechTree(techVal, depth))
		}
	}
	return data
}

func GetSimilar(c *gin.Context) {
	uuid := c.Param("uuid")
	page, _ := strconv.Atoi(c.Param("page"))
	fund, err := dbops.GetFundDetail(uuid)
	if err != nil {
		log.Error("dbops GetFundDetail error: %s", err)
		sendErrorResponse(c, defs.ErrorDBError)
		return
	}
	dbSearch := &dbops.SearchParams{
		From: strconv.Itoa(page * 10),
	}
	orgs := []string{}
	for _, element := range fund.Institution {
		orgs = append(orgs, element.Name)
	}
	org := strings.Join(orgs, " ")
	tech := strings.Join(fund.Technology, " ")
	industries := strings.Join(fund.Industries, " ")
	dbSearch.Organization = org
	dbSearch.Technology = tech
	dbSearch.Industries = industries
	elements, err := dbops.GetMultiSearch(dbSearch)
	if err != nil {
		log.Error("multisearch error: %s", err)
		sendErrorResponse(c, defs.ErrorResponse{
			HTTPSc: http.StatusInternalServerError,
			Error: defs.Err{
				Error:     fmt.Sprintf("%s", err),
				ErrorCode: elasticsearchError,
			},
		})
		return
	}
	keywords_bucket := append(orgs, fund.Technology...)
	keywords_bucket = append(keywords_bucket, fund.Industries...)
	number := 10
	if len(elements) < 10 {
		number = len(elements)
	}
	data := make([]*defs.SimilarData, number)
	for i := 0; i < number; i++ {
		element := elements[i]
		data[i] = &defs.SimilarData{
			Title: element.AwardTitle,
			UUID:  element.AwardID,
		}
		keywords := make([]string, 3)
		ids := generateRandomNumber(0, len(keywords_bucket)-1, 3)
		for i, id := range ids {
			keywords[i] = keywords_bucket[id]
		}
		data[i].Keywords = keywords
	}
	sendNormalResponse(c, defs.NormalResp{
		HttpSc: http.StatusOK,
		Resp: &defs.SimilarDataResponse{
			Data: data,
		},
	})
	return
}

func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
