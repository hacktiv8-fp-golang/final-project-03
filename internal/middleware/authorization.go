package middleware

import (
	"github.com/hacktiv8-fp-golang/final-project-03/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-03/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		userData := context.MustGet("userData").(jwt.MapClaims)
		userRole := userData["role"].(string)

		if userRole != "admin" {
			err := helper.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func CategoryAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		categoryId, err := helper.GetIdParam(context, "categoryId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		db := database.GetDB()
		category := model.Category{}

		errMsg := db.First(&category, categoryId).Error
		if errMsg != nil {
			err := helper.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func TaskAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		taskId, err := helper.GetIdParam(context, "taskId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		task := model.Task{}

		errMsg := db.Select("user_id").First(&task, taskId).Error
		if errMsg != nil {
			err := helper.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if task.UserID != userID {
			err := helper.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
