package scanner

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/piggona/fundings_view/scheduler/crontask"
	"github.com/piggona/fundings_view/scheduler/middleware"
	"github.com/piggona/fundings_view/scheduler/models"
	"github.com/piggona/fundings_view/scheduler/utils/log"
)

var (
	NoDataDispatched = fmt.Errorf("dispatcher got empty data, finished")
)

type Investigator struct {
	FirstName    string `xml:"FirstName" json:"firstname"`
	LastName     string `xml:"LastName" json:"lastname"`
	EmailAddress string `xml:"EmailAddress" json:"email_address"`
	StartDate    string `xml:"StartDate" json:"start_date"`
	EndDate      string `xml:"EndDate" json:"end_date"`
	RoleCode     string `xml:"RoleCode" json:"role_code"`
}

type Institution struct {
	Name          string `xml:"Name" json:"name"`
	CityName      string `xml:"CityName" json:"cityname"`
	ZipCode       string `xml:"ZipCode" json:"zipcode"`
	PhoneNumber   string `xml:"PhoneNumber" json:"phone_number"`
	StreetAddress string `xml:"StreetAddress" json:"street_address"`
	CountryName   string `xml:"CountryName" json:"country_name"`
	StateName     string `xml:"StateName" json:"state_name"`
	StateCode     string `xml:"StateCode" json:"state_code"`
}

type Organization struct {
	Code        string `xml:"Code" json:"code"`
	Directorate string `xml:"Directorate>LongName" json:"directorate"`
	Division    string `xml:"Division>LongName" json:"division"`
}

type ProgramElement struct {
	Code string `xml:"Code" json:"code"`
	Text string `xml:"Text" json:"text"`
}

type ProgramReference struct {
	Code string `xml:"Code" json:"code"`
	Text string `xml:"Text" json:"text"`
}

type FundContent struct {
	XMLName             xml.Name            `xml:"Award"`
	AwardID             string              `xml:"AwardID" json:"award_id"`
	AwardTitle          string              `xml:"AwardTitle" json:"award_title"`
	AwardEffectiveDate  string              `xml:"AwardEffectiveDate" json:"award_effective_date"`
	AwardExpirationDate string              `xml:"AwardExpirationDate" json:"award_expiration_date"`
	AwardAmount         float32             `xml:"AwardAmount" json:"award_amount"`
	Description         string              `xml:"AbstractNarration" json:"description"`
	Investigators       []*Investigator     `xml:"Investigator" json:"investigator"`
	Institutions        []*Institution      `xml:"Institution" json:"institution"`
	Organizations       []*Organization     `xml:"Organization" json:"organization"`
	ProgramElements     []*ProgramElement   `xml:"ProgramElement" json:"program_element"`
	ProgramReferences   []*ProgramReference `xml:"ProgramReference" json:"program_reference"`
}

type Fund struct {
	XMLName xml.Name     `xml:"rootTag"`
	Award   *FundContent `xml:"Award"`
}

type FileTask struct {
	id   int
	file string
}

type Scanner struct {
	offset int
}

func NewScanner() *Scanner {
	return &Scanner{
		offset: 0,
	}
}

func (s *Scanner) Dispatcher(dc crontask.DataChan) error {
	importTask, err := models.GetImportTask(s.offset)
	if err != nil {
		if err == orm.ErrNoRows {
			return crontask.FinishErr
		}
		log.Error("error in get import task, no data got")
		return err
	}
	err = models.SetImportTaskStatus(importTask.Id, models.STATUS_RUNNING)
	if err != nil {
		log.Error("error in setimporttaskstatus: %s", err)
		return err
	}
	files, err := ioutil.ReadDir(importTask.Path)
	if err != nil {
		log.Error("read dir error: %s", err)
		return err
	}
	go func() {
		for _, file := range files {
			if !file.IsDir() {
				dc <- &FileTask{
					id:   importTask.Id,
					file: importTask.Path + string(os.PathSeparator) + file.Name(),
				}
			}
		}
		dc <- "FINISH"
	}()
	s.offset++
	return nil
}

func (s *Scanner) Executor(dc crontask.DataChan) error {
	errMap := &sync.Map{}
	var err error
	var wg sync.WaitGroup
	var taskid int
	token := make(chan struct{}, 50)
forloop:
	for {
		select {
		case task := <-dc:
			if _, ok := task.(string); ok {
				break forloop
			}
			wg.Add(1)
			token <- struct{}{}
			go func(file interface{}) {
				defer func() {
					<-token
					wg.Done()
				}()
				filetask := task.(*FileTask)
				taskid = filetask.id
				fund := new(Fund)
				xmlContent, err := ioutil.ReadFile(filetask.file)
				if err != nil {
					log.Error("read xml file error: %s", err)
					return
				}
				err = xml.Unmarshal(xmlContent, fund)
				if err != nil {
					log.Error("xml unmarshal error: %s", err)
					return
				}
				bytes, err := json.Marshal(fund.Award)
				if err != nil {
					log.Error("json marshal error: %s", err)
					return
				}
				middleware.PutData("import_raw", string(bytes))
			}(task)
		}
	}
	wg.Wait()
	models.SetImportTaskStatus(taskid, models.STATUS_FINISH)
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}

func (s *Scanner) Reset() {
	s.offset = 0
}
