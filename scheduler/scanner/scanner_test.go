package scanner

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/piggona/fundings_view/scheduler/middleware"
	"github.com/piggona/fundings_view/scheduler/utils/log"
)

func TestMain(m *testing.M) {
	brokers := []string{"172.30.39.100:9092"}
	middleware.InitProducer(context.Background(), brokers)
	m.Run()
}

func TestScanner(t *testing.T) {
	fund := new(Fund)
	xmlContent, err := ioutil.ReadFile("./1900008.xml")
	if err != nil {
		log.Error("read xml file error: %s", err)
		return
	}
	err = xml.Unmarshal(xmlContent, fund)
	if err != nil {
		log.Error("xml unmarshal error: %s", err)
		return
	}
	log.Info("%v", fund)
	bytes, err := json.Marshal(fund.Award)
	if err != nil {
		log.Error("json marshal error: %s", err)
		return
	}
	middleware.PutData("import_raw", string(bytes))

}

func TestFetchData(t *testing.T) {

}
