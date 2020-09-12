package models

import (
	"github.com/astaxie/beego/orm"
)

const (
	STATUS_READY   = "r"
	STATUS_RUNNING = "i"
	STATUS_FINISH  = "f"
)

type ImportTask struct {
	Id     int `orm:"pk"`
	Path   string
	Type   string
	Status string
}

func init() {
	orm.RegisterModel(new(ImportTask))
}

func GetImportTask(offset int) (importTask *ImportTask, err error) {
	tasks := []*ImportTask{}
	_, err = orm.NewOrm().QueryTable(new(ImportTask)).Filter("status", STATUS_READY).Filter("type", "0").Limit(1, offset).All(&tasks)
	if err != nil {
		return
	}
	if len(tasks) == 0 {
		return nil, orm.ErrNoRows
	}
	importTask = tasks[0]
	return
}

func SetImportTaskStatus(id int, status string) (err error) {
	_, err = orm.NewOrm().Update(&ImportTask{Id: id, Status: status}, "status")
	return
}
