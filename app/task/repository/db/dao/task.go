package dao

import (
	"context"
	"todo_list/app/task/repository/db/model"
	"todo_list/idl/pb"

	"gorm.io/gorm"
)

// 定义的是对数据库的 Task model的 CRUD 操作

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}
func (dao *TaskDao) CreateTask(data *model.Task) (err error) {
	return dao.Model(&model.Task{}).Create(&data).Error
}

// ListTaskByUserId 获取user id
func (dao *TaskDao) ListTaskByUserId(userID uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).
		Limit(limit).Where("uid = ?", userID).
		Find(&r).Error

	if err != nil {
		return
	}
	err = dao.Model(&model.Task{}).Where("uid = ?", userID).
		Count(&count).Error
	return
}

// GetTaskByTaskIdAndUserId 通过 task id 和 user id 获取task
func (dao *TaskDao) GetTaskByTaskIdAndUserId(tId, uId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", tId, uId).
		Find(&r).Error
	return
}
func (dao *TaskDao) DeleteTaskByIdAndUserId(tId, uId uint64) (err error) {
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", tId, uId).
		Find(&model.Task{}).Error
	return
}

// UpdateTask 更新task
func (dao *TaskDao) UpdateTask(req *pb.TaskRequest) (r *model.Task, err error) {
	r = new(model.Task)
	err = dao.Model(&model.Task{}).
		Where("id = ? AND uid = ?", req.Id, req.Uid).
		First(&r).Error
	if err != nil {
		return
	}
	r.Title = req.Title
	r.Status = int(req.Status)
	r.Content = req.Content

	err = dao.Save(&r).Error
	return
}
