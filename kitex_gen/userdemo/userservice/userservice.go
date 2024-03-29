// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservice

import (
	"context"
	userdemo "ez-note/kitex_gen/userdemo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userdemo.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newUserServiceMGetUserArgs, newUserServiceMGetUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "userdemo",
		"ServiceFilePath": `idl/user.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdemo.UserServiceCreateUserArgs)
	realResult := result.(*userdemo.UserServiceCreateUserResult)
	success, err := handler.(userdemo.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return userdemo.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return userdemo.NewUserServiceCreateUserResult()
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdemo.UserServiceMGetUserArgs)
	realResult := result.(*userdemo.UserServiceMGetUserResult)
	success, err := handler.(userdemo.UserService).MGetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceMGetUserArgs() interface{} {
	return userdemo.NewUserServiceMGetUserArgs()
}

func newUserServiceMGetUserResult() interface{} {
	return userdemo.NewUserServiceMGetUserResult()
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*userdemo.UserServiceCheckUserArgs)
	realResult := result.(*userdemo.UserServiceCheckUserResult)
	success, err := handler.(userdemo.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return userdemo.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return userdemo.NewUserServiceCheckUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (r *userdemo.CreateUserResponse, err error) {
	var _args userdemo.UserServiceCreateUserArgs
	_args.Req = req
	var _result userdemo.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (r *userdemo.MGetUserResponse, err error) {
	var _args userdemo.UserServiceMGetUserArgs
	_args.Req = req
	var _result userdemo.UserServiceMGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (r *userdemo.CheckUserResponse, err error) {
	var _args userdemo.UserServiceCheckUserArgs
	_args.Req = req
	var _result userdemo.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
