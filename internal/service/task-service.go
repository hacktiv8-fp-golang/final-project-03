package service

import (
	"github.com/hacktiv8-fp-golang/final-project-03/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/repository"

	"github.com/asaskevich/govalidator"
)

type taskServiceRepo interface {
	CreateTask(task *model.Task, userID uint) (*model.Task, helper.Error)
	GetAllTasks() ([]*model.Task, helper.Error)
	UpdateTask(task *model.TaskUpdate, taskId uint) (*model.Task, helper.Error)
	UpdateStatusTask(statusupdate *model.TaskStatusUpdate, taskId uint) (*model.Task, helper.Error)
	UpdateCategoryIdTask(Updatetask *model.TaskCategoryUpdate, taskId uint) (*model.Task, helper.Error)
	DeleteTask(taskId uint) helper.Error
}

type taskService struct{}

var TaskService taskServiceRepo = &taskService{}

func (t *taskService) CreateTask(task *model.Task, userID uint) (*model.Task, helper.Error){
	task.UserID = userID

	if _, err := govalidator.ValidateStruct(task); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	_, err := repository.CategoryModel.GetCategoryById(task.CategoryID);

	if err != nil {
		return nil, err
	}

	createdTask, err := repository.TaskModel.CreateTask(task)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (t *taskService) GetAllTasks() ([]*model.Task, helper.Error) {
	var task []*model.Task

	task, err := repository.TaskModel.GetAllTasks()

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) UpdateTask(task *model.TaskUpdate, taskId uint) (*model.Task, helper.Error) {
	if _, err := govalidator.ValidateStruct(task); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.TaskModel.UpdateTask(task, taskId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *taskService) UpdateStatusTask(statusupdate *model.TaskStatusUpdate, taskId uint) (*model.Task, helper.Error) {
	if _, err := govalidator.ValidateStruct(statusupdate); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.TaskModel.UpdateStatusTask(statusupdate, taskId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *taskService) UpdateCategoryIdTask(Updatetask *model.TaskCategoryUpdate, taskId uint) (*model.Task, helper.Error) {
	if _, err := govalidator.ValidateStruct(Updatetask); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	_, err := repository.CategoryModel.GetCategoryById(Updatetask.CategoryID);

	if err != nil {
		return nil, err
	}

	result, err := repository.TaskModel.UpdateCategoryIdTask(Updatetask, taskId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *taskService) DeleteTask(taskId uint) helper.Error {
	err := repository.TaskModel.DeleteTask(taskId)

	if err != nil {
		return err
	}

	return nil
}
