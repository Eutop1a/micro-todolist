package http

import (
	"net/http"
	"todo_list/app/gateway/rpc"
	"todo_list/idl/pb"
	"todo_list/pkg/ctl"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBindJSON"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-TaskCreate"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UpdateTaskHandler-ShouldBindJSON"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UpdateTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UpdateTaskHandler-TaskUpdate"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func ListTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "ListTaskHandler-ShouldBindJSON"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "ListTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "ListTaskHandler-TaskList"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "DeleteTaskHandler-ShouldBindJSON"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "DeleteTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "DeleteTaskHandler-TaskDelete"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskHandler-ShouldBindJSON"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskGet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskHandler-TaskGet"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}
