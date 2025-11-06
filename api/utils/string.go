package utils

import (
	"compus-second-hand/global"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 获取mysql的地址
func GetDbAddress() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.Configs.MySQL.User, global.Configs.MySQL.Password, global.Configs.MySQL.Host, global.Configs.MySQL.Port, global.Configs.MySQL.Name)
}

// md5加密
func Md5(str string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(str))
	if err != nil {
		return "", err
	}
	md5String := hex.EncodeToString(h.Sum(nil))
	return md5String, nil
}
