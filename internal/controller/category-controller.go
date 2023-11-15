package controller

import (
	"final-project-03/internal/helper"
	"final-project-03/internal/model"
	"final-project-03/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Creates a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param model.Category body model.CategoryCreate true "Category object to be created"
// @Success 201 {object} model.User "Category created successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /categories [post]
func CreateCategory(context *gin.Context) {
	var category model.Category

	if err := context.ShouldBindJSON(&category); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.CategoryService.CreateCategory(&category)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         result.ID,
		"type":       result.Type,
		"created_at": result.CreatedAt,
	})
}

// UpdateCategory godoc
// @Summary Update a Category.
// @Tags Categories
// @Accept json
// @Produce json
// @Param categoryId path int true "Category ID"
// @Param model.Category body model.CategoryUpdate true "Category object to be updated"
// @Success 200 {object} model.User "Category updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /categories/{categoryId} [put]
func UpdateCategory(context *gin.Context) {
	var update model.CategoryUpdate

	if err := context.ShouldBindJSON(&update); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	categoryId, _ := helper.GetIdParam(context, "categoryId")

	result, err := service.CategoryService.UpdateCategory(&update, categoryId)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"type":       result.Type,
		"updated_at": result.UpdatedAt,
	})
}
