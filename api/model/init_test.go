package model

import (
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	"fmt"
	"testing"
)

func TestDBinit(t *testing.T) {
	utils.ConfigInit()
	DBinit()
	var users = &User{
		Username: "admin1",
		Gender:   0,
		Password: "123456",
		Email:    "admin@ewwxample.com",
		CampusID: 1,
	}
	res, err := global.DB.NamedExec("insert into user (username, password, gender, email, campus_id) values (:username, :password, :gender, :email, :campus_id)", users)
	if err != nil {
		t.Fatalf("插入用户失败: %v", err)
	}
	fmt.Println(res.LastInsertId())
}
