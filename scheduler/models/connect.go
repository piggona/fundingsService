package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(passwd, host string) {
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		panic(err)
	}

	connectStr := fmt.Sprintf("root:%s@(%s)/fundingsview?charset=utf8&timeout=5s", passwd, host)
	if err := orm.RegisterDataBase("default", "mysql", connectStr); err != nil {
		panic(err)
	}

	orm.SetMaxOpenConns("default", 30)

	if defaultDB, err := orm.GetDB("fundingsview"); err == nil {
		defaultDB.SetConnMaxLifetime(3395 * time.Second)
		err = defaultDB.Ping()
		if err != nil {
			panic(err)
		}
	}
	orm.RegisterModel(new(ImportTask))
}
