package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) ListTasks(c *gin.Context) {
	tasks, err := s.TaskDAO.List(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}
