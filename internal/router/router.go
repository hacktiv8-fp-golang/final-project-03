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
		userRouter.PUT("/update-account", middleware.Authentication(), controller.UpdateUser)
		userRouter.DELETE("/delete-account", middleware.Authentication(), controller.DeleteUser)
	}

	categoryRouter := router.Group("/categories")
	{
		categoryRouter.Use(middleware.Authentication())
		categoryRouter.POST("/", middleware.AdminAuthorization(), controller.CreateCategory)
		categoryRouter.GET("/", controller.GetAllCategories)
		categoryRouter.PATCH("/:categoryId", middleware.AdminAuthorization(), middleware.CategoryAuthorization(), controller.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", middleware.AdminAuthorization(), middleware.CategoryAuthorization(), controller.DeleteCategory)
	}

	taskRouter := router.Group("/tasks")
	{
		taskRouter.Use(middleware.Authentication())
		taskRouter.POST("/", controller.CreateTask)
		taskRouter.GET("/", controller.GetAllTasks)
		taskRouter.PUT("/:taskId", middleware.TaskAuthorization(), controller.UpdateTask)
		taskRouter.PATCH("/update-status/:taskId", middleware.TaskAuthorization(), controller.UpdateStatusTask)
		taskRouter.PATCH("/update-category/:taskId", middleware.TaskAuthorization(), controller.UpdateCategoryIdTask)
		taskRouter.DELETE("/:taskId", middleware.TaskAuthorization(), controller.DeleteTask)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}
