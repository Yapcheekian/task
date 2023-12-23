package main

import (
	"net/http"

	"github.com/Yapcheekian/task/pkg/dao"
	"github.com/Yapcheekian/task/pkg/model"
	"github.com/Yapcheekian/task/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	taskDAO := &dao.MemoryTaskDAO{
		Store: make([]*model.Task, 0),
	}

	svc := service.Service{
		TaskDAO: taskDAO,
	}

	r := gin.Default()
	r.GET("/tasks", svc.ListTasks)
	r.POST("/tasks", svc.CreateTask)
	r.PUT("/tasks/:id", svc.UpdateTask)
	r.DELETE("/tasks/:id", svc.DeleteTask)

	// health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
