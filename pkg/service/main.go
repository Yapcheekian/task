package service

import (
	"github.com/Yapcheekian/task/pkg/dao"
	"github.com/gin-gonic/gin"
)

type Service struct {
	TaskDAO dao.TaskDAO
}

func errorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}
