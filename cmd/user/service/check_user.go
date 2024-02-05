package service

import (
	"context"
	"crypto/md5"
	"ez-note/cmd/user/dal/db"
	"ez-note/kitex_gen/userdemo"
	"ez-note/pkg/errno"
	"fmt"
	"io"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *userdemo.CheckUserRequest) (int64, error) {
	// 1密码加密
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		// 2检查有无
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	// 3验证密码
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	// 4返回ID，用于生成Token
	return int64(u.ID), nil
}
