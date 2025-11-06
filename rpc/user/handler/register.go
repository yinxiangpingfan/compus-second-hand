package handler

import (
	"compus-second-hand/api/model"
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	user "compus-second-hand/rpc/user/pb"
	"context"
)

func (this *UserHandlerEr) Register(ctx context.Context, req *user.RegisterRequest, resp *user.RegisterResponse) error {
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
		Gender:   int(req.Gender),
		CampusID: uint64(req.Campus),
		Pic:      req.File,
	}
	res, err := global.DB.NamedExec("insert into user (email, password, username, gender, campus_id, pic) values (:email, :password, :username, :gender, :campus_id, :pic)", user)
	if err != nil {
		codeErr := utils.CheckError(err)
		switch codeErr {
		case 1:
			resp.Code = 1 // 唯一键冲突
		case 2:
			resp.Code = 2 // 外键约束冲突
		case 3:
			resp.Code = 3 // 空值约束冲突
		default:
			return err // 未知错误
		}
	}
	id, err := res.RowsAffected()
	if err != nil {
		// 插入失败
		return err
	}
	if id == 0 {
		resp.Code = 4 // 插入失败
		return nil
	}
	resp.Code = 0 // 插入成功
	return nil
}
