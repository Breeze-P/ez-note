package handlers

import (
	"context"
	"ez-note/cmd/api/rpc"
	"ez-note/kitex_gen/userdemo"
	"ez-note/pkg/errno"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
)

// Register register user info
func Register(ctx context.Context, c *app.RequestContext) {
	// 1绑定参数
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	log.Printf("%#v\n", registerVar)
	log.Printf("%#v\n", c.PostForm("username"))
	// 合法性判断
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	// rpc调用
	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
