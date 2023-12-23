package service

import (
	"net/http"
	"strconv"

	"github.com/Yapcheekian/task/pkg/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	formattedTaskID, err := strconv.Atoi(taskID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	task.ID = formattedTaskID

	if err := s.TaskDAO.Update(c.Request.Context(), &task); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
