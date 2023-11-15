package controller

import (
	"final-project-03/internal/helper"
	"final-project-03/internal/model"
	"final-project-03/internal/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register an Account.
// @Tags Users
// @Accept json
// @Produce json
// @Param model.User body model.UserCreate true "User object to be created"
// @Success 201 {object} model.User "User created successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /users/register [post]
func Register(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Register(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         result.ID,
		"full_name":  result.FullName,
		"email":      result.Email,
		"created_at": result.CreatedAt,
	})
}

// Login godoc
// @Summary Login an Account.
// @Tags Users
// @Accept json
// @Produce json
// @Param model.User body model.LoginCredential true "User login credentials"
// @Success 200 {object} model.User "User logged in successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /users/login [post]
func Login(context *gin.Context) {
	var user model.LoginCredential

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Login(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": result})
}

// UpdateAccount godoc
// @Summary Update an Account.
// @Tags Users
// @Accept json
// @Produce json
// @Param model.User body model.UserUpdate true "User object to be updated"
// @Success 200 {object} model.User "User updated successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /users/update-account [put]
func UpdateUser(context *gin.Context) {
	var update model.UserUpdate

	if err := context.ShouldBindJSON(&update); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	result, err := service.UserService.UpdateUser(userID, &update)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"email":      result.Email,
		"full_name":  result.FullName,
		"updated_at": result.UpdatedAt,
	})
}

// DeleteAccount godoc
// @Summary Delete an Account.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.User "User deleted successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 422 {object} helper.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} helper.ErrorResponse "Server Error"
// @Router /users/delete-account [delete]
func DeleteUser(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	_, err := service.UserService.DeleteUser(userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your Account has been successfully deleted",
	})
}
