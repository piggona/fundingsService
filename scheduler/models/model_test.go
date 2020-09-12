package models

import (
	"testing"

	"github.com/piggona/fundings_view/scheduler/utils/log"
)

func TestGetImportTask(t *testing.T) {
	importTask, err := GetImportTask(0)
	if err != nil {
		t.Error(err)
	}
	log.Info("%v", importTask)
}

func TestSetImportTaskStatus(t *testing.T) {
	SetImportTaskStatus(1, STATUS_READY)
}
