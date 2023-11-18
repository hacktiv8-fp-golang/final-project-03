package router

import (
	"final-project-03/internal/controller"
	"final-project-03/internal/middleware"

	"github.com/gin-gonic/gin"

	_ "final-project-03/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.PUT("/update-account", controller.UpdateUser)
		userRouter.DELETE("/delete-account", controller.DeleteUser)
	}

	categoryRouter := router.Group("/categories")
	{
		categoryRouter.Use(middleware.Authentication())
		categoryRouter.POST("/", middleware.AdminAuthorization(), controller.CreateCategory)
		categoryRouter.GET("/", controller.GetAllCategories)
		categoryRouter.PATCH("/:categoryId",  middleware.AdminAuthorization(), middleware.CategoryAuthorization(), controller.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", middleware.AdminAuthorization(), middleware.CategoryAuthorization(), controller.DeleteCategory)
	}

	taskRouter := router.Group("/comments")
	{
		taskRouter.POST("/")
		taskRouter.GET("/")
		taskRouter.PUT("/:taskId")
		taskRouter.PATCH("/update-status/:taskId")
		taskRouter.DELETE("/:taskId")
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}
