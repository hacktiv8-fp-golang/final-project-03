package middleware

import (
	"final-project-03/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifiedToken, err  := helper.VerifyToken(context)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.Set("userData", verifiedToken)
		context.Next()
	}
}
