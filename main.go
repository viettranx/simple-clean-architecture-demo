package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-clean-architecture-demo/modules/task/business"
	"simple-clean-architecture-demo/modules/task/repository/inmem"
	"simple-clean-architecture-demo/modules/task/transport/rest"
)

func main() {
	engine := gin.Default()

	// Setup full service dependencies
	apiService := rest.NewAPI(business.NewBusiness(inmem.NewInMemStorage()))

	v1 := engine.Group("v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", apiService.CreateTaskHdl())
			tasks.GET("", apiService.ListTaskHdl())
			tasks.GET("/:id", apiService.GetTaskHdl())
			tasks.PATCH("/:id", apiService.UpdateTaskHdl())
			tasks.DELETE("/:id", apiService.DeleteTaskHdl())
		}
	}

	if err := engine.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
