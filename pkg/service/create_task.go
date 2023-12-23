package service

import (
	"net/http"

	"github.com/Yapcheekian/task/pkg/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) CreateTask(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := s.TaskDAO.Create(c.Request.Context(), &task); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
