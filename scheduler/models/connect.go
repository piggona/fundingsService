package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		panic(err)
	}

	if err := orm.RegisterDataBase("default", "mysql", "root:vanlink@(172.30.39.100:3306)/fundingsview?charset=utf8&timeout=5s"); err != nil {
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
}
