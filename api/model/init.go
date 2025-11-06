package model

import (
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DBinit() {
	var err error
	//获取数据库的地址
	dsn := utils.GetDbAddress()
	global.DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	fmt.Println("数据库连接成功")
}
