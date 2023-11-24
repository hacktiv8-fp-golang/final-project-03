package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-03/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary Creates a new task
// @Tags Task
// @Accept json
// @Produce json
// @Param model.Category body model.TaskCreate true "Task object to be created"
// @Success 201 {object} model.TaskCreateResponse "Task created successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /tasks [post]
func CreateTask(context *gin.Context){
	var task model.Task

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
    context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
    return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	task.Status = false


	createdResponse, err := service.TaskService.CreateTask(&task, userID)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id": 					createdResponse.ID,
		"title": 				createdResponse.Title,
		"status": 			createdResponse.Status,
		"description": 	createdResponse.Description,
		"user_id": 			createdResponse.UserID,
		"category_id":  createdResponse.CategoryID,
		"created_at": 	createdResponse.CreatedAt,
	})
}

// GetAllTasks godoc
// @Summary Get all task
// @Description Retrieve a list of all task
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} []model.Task
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /tasks [get]
func GetAllTasks (context *gin.Context){

	results, err := service.TaskService.GetAllTasks()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	tasks := make([]gin.H, 0, len(results))

	for _, result := range results {
		task := gin.H{
			"id": 					result.ID,
			"title": 				result.Title,
			"status": 			result.Status,
			"description": 	result.Description,
			"user_id": 			result.UserID,
			"category_id":  result.CategoryID,
			"created_at": 	result.CreatedAt,
			"User": gin.H{
				"id":       	result.User.ID,
				"email":    	result.User.Email,
				"full_name": 	result.User.FullName,
			},
		}

		tasks = append(tasks, task)
	}

	context.JSON(http.StatusOK, tasks)
}

// UpdateTask godoc
// @Summary Update a Task.
// @Tags Task
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param model.Task body model.TaskUpdate true "Task object to be updated"
// @Success 200 {object} model.TaskUpdateResponse "Task updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /tasks/{taskId} [put]
func UpdateTask (context *gin.Context){
	var task model.TaskUpdate

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	taskIDInt, _ := strconv.Atoi(context.Param("taskId"))
	taskID := uint(taskIDInt)

	result, err := service.TaskService.UpdateTask(&task,taskID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         	result.ID,
		"title":      	result.Title,
		"description":  result.Description,
		"status":  			result.Status,
		"user_id":    	result.UserID,
		"category_id":	result.CategoryID,
		"updated_at": 	result.CreatedAt,
	})
}

// UpdateTaskStatus godoc
// @Summary Update a Task.
// @Tags Task
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param model.Task body model.TaskStatusUpdate true "Task Status object to be updated"
// @Success 200 {object} model.TaskUpdateResponse "Task Status updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /tasks/update-status/{taskId} [patch]
func UpdateStatusTask (context *gin.Context){
	var task model.TaskStatusUpdate

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	taskIDInt, _ := strconv.Atoi(context.Param("taskId"))
	taskID := uint(taskIDInt)

	result, err := service.TaskService.UpdateStatusTask(&task,taskID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         	result.ID,
		"title":      	result.Title,
		"description":  result.Description,
		"status":  			result.Status,
		"user_id":    	result.UserID,
		"category_id":	result.CategoryID,
		"updated_at": 	result.CreatedAt,
	})
}

// UpdateTaskCategory godoc
// @Summary Update a Task.
// @Tags Task
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param model.Task body model.TaskCategoryUpdate true "Task Category object to be updated"
// @Success 200 {object} model.TaskUpdateResponse "Task Category updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /tasks/update-category/{taskId} [patch]
func UpdateCategoryIdTask (context *gin.Context){
	var task model.TaskCategoryUpdate

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	taskIDInt, _ := strconv.Atoi(context.Param("taskId"))
	taskID := uint(taskIDInt)

	result, err := service.TaskService.UpdateCategoryIdTask(&task, taskID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         	result.ID,
		"title":      	result.Title,
		"description":  result.Description,
		"status":  			result.Status,
		"user_id":    	result.UserID,
		"category_id":	result.CategoryID,
		"updated_at": 	result.CreatedAt,
	})
}

// DeleteTask godoc
// @Summary Delete a task item
// @Description Delete a task item by id
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param taskId path int true "Task id"
// @Success 200 {string} string "OK"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Security Bearer
// @Router /task/{taskId} [delete]
func DeleteTask (context *gin.Context){
	id, _ := helper.GetIdParam(context, "taskId")

	err := service.TaskService.DeleteTask(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Task has been successfully deleted",
	})
}
