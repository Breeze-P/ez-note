package service

import (
	"context"
	"ez-note/cmd/user/dal/db"
	"ez-note/cmd/user/pack"
	"ez-note/kitex_gen/userdemo"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
