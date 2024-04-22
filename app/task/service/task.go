package service

import (
	"context"
	"encoding/json"
	"sync"
	"todo_list/app/task/repository/db/dao"
	"todo_list/app/task/repository/db/model"
	"todo_list/app/task/repository/mq"
	"todo_list/idl/pb"
	"todo_list/pkg/e"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

// GetTaskSrv 懒汉式的单例模式 lazy-loading
func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

// CreateTask create task 送到mq, mq --> 落库，
func (t *TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	body, _ := json.Marshal(req)
	resp.Code = e.Success
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func TaskMQ2DB(ctx context.Context, req *pb.TaskRequest) error {
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(m)
}

func (t *TaskSrv) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = e.Success
	if req.Limit == 0 {
		req.Limit = 10
	}
	r, count, err := dao.NewTaskDao(ctx).ListTaskByUserId(req.Uid, int(req.StartTime), int(req.Limit))
	if err != nil {
		resp.Code = e.Error
		return
	}
	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)

	return
}

func (t *TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	r, err := dao.NewTaskDao(ctx).GetTaskByTaskIdAndUserId(req.Id, req.Uid)
	if r.ID == 0 || err != nil {
		resp.Code = e.Error
		return
	}
	resp.TaskDetail = BuildTask(r)

	return
}

func (t *TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.Error
		return
	}

	return
}

func (t *TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).DeleteTaskByTaskIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = e.Error
		return
	}

	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	taskModel := pb.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}
