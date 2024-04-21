package http

import (
	"net/http"
	"todo_list/app/gateway/rpc"
	"todo_list/idl/pb"
	"todo_list/pkg/ctl"
	"todo_list/pkg/jwt"
	"todo_list/types"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserRegister-RPC"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
}
func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-ShouldBind"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-UserRegister-RPC"))
		return
	}

	token, err := jwt.GenerateToken(uint(userResp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-GenerateToken"))
		return
	}
	res := &types.TokenData{
		User:  userResp,
		Token: token,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
