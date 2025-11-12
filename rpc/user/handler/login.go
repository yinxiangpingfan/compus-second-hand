package handler

import (
	"compus-second-hand/api/model"
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	user "compus-second-hand/rpc/user/pb"
	"context"
)

//code = 1，密码错误 code = 0，登陆成功 code = 2，用户不存在

func (*UserHandlerEr) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) error {
	sqlStr := "select id, username, password from user where "
	var user model.User
	var boolUser bool = false //用于记录是否查询到用户
	if req.Email == "" {
		//没有邮箱的情况下从数据库查用户名获取信息
		sqlStr += "username = :username"
		rows, err := global.DB.NamedQuery(sqlStr, map[string]interface{}{"username": req.Username})
		if err != nil {
			return err
		}
		boolUser = rows.Next()
		if boolUser {
			err := rows.Scan(&user.ID, &user.Username, &user.Password)
			if err != nil {
				return err
			}
		}
	} else {
		//有邮箱的情况下从数据库查邮箱获取信息
		sqlStr += "email = :email"
		rows, _ := global.DB.NamedQuery(sqlStr, map[string]interface{}{"email": req.Email})
		boolUser = rows.Next()
		if boolUser {
			err := rows.Scan(&user.ID, &user.Username, &user.Password)
			if err != nil {
				return err
			}
		}
	}
	if !boolUser {
		resp.Code = 2
		return nil
	}
	md5Pass, err := utils.Md5(req.Password)
	if err != nil {
		return err
	}
	if md5Pass != user.Password {
		resp.Code = 1
	}
	resp.Code = 0
	resp.Id = int64(user.ID)
	return nil
}
