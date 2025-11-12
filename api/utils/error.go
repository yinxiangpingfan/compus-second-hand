package utils

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

const (
	DBErrUnique  uint16 = 1062 //唯一键冲突
	DBErrFK      uint16 = 1452 //外键约束冲突
	DBErrBadNull uint16 = 1048 //空值约束冲突
)

func CheckError(err error) int {
	var e *mysql.MySQLError
	if errors.As(err, &e) {
		switch e.Number {
		case DBErrUnique:
			return 1 // 唯一键冲突
		case DBErrFK:
			return 2 //外键约束冲突
		case DBErrBadNull:
			return 3 //空值约束冲突
		}
	}
	return 4 // 未知错误
}
