package rpc

import (
	"context"
	"todo_list/idl/pb"
	"todo_list/pkg/e"
)

func TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.CreateTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}

func TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.UpdateTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}

func TaskList(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	resp, err = TaskService.GetTasksList(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}

func TaskGet(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.GetTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}

func TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.DeleteTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}
