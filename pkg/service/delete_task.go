package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Service) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	formattedTaskID, err := strconv.Atoi(taskID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := s.TaskDAO.Delete(c.Request.Context(), formattedTaskID); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
