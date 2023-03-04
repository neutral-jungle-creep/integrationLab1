package v1

import (
	"IntegrationLab1/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type UserRequest struct {
	FirstName string `json:"first-name" binding:"required"`
	LastName  string `json:"last-name" binding:"required"`
}

func (h *Handler) SubmitCompletedDoc(c *gin.Context) {
	var input UserRequest

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ответ": "данные получены успешно",
	})
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, message)
}
