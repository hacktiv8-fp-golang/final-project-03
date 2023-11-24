package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-03/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/model"
)

type taskModelRepo interface {
	CreateTask(task *model.Task) (*model.Task, helper.Error)
	GetAllTasks() ([]*model.Task, helper.Error)
	UpdateTask(input *model.TaskUpdate, taskID uint) (*model.Task, helper.Error)
	UpdateStatusTask(taskStatus *model.TaskStatusUpdate, taskID uint) (*model.Task, helper.Error)
	UpdateCategoryIdTask(input *model.TaskCategoryUpdate, taskID uint) (*model.Task, helper.Error)

	DeleteTask(taskId uint) helper.Error
}

type taskModel struct{}

var TaskModel taskModelRepo = &taskModel{}

func (t *taskModel) CreateTask(task *model.Task) (*model.Task, helper.Error){
	db := database.GetDB()

	err:= db.Create(&task).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return task, nil
}

func (t *taskModel) GetAllTasks() ([]*model.Task, helper.Error) {
	db := database.GetDB()
	var tasks []*model.Task

	err := db.Preload("User").Find(&tasks).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return tasks, nil
}

func (t *taskModel) UpdateTask(input *model.TaskUpdate, taskID uint) (*model.Task, helper.Error) {
	db := database.GetDB()

	var task model.Task
	err := db.First(&task, taskID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&task).Updates(input)

	return &task, nil
}

func (t *taskModel) UpdateStatusTask(taskStatus *model.TaskStatusUpdate, taskID uint) (*model.Task, helper.Error) {
	db := database.GetDB()

	var task model.Task

	err := db.First(&task, taskID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&task).Updates(taskStatus)

	return &task, nil
}

func (t *taskModel) UpdateCategoryIdTask(input *model.TaskCategoryUpdate, taskID uint) (*model.Task, helper.Error) {
	db := database.GetDB()

	var task model.Task

	err := db.First(&task, taskID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&task).Updates(input)

	return &task, nil
}

func (t *taskModel) DeleteTask(taskId uint) helper.Error {
	db := database.GetDB()

	var task model.Task

	err := db.Where("id = ?", taskId).Delete(&task).Error

	if err != nil {
		return helper.ParseError(err)
	}

	return nil
}
