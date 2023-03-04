package v1

import (
	"IntegrationLab1/internal/domain"
	"IntegrationLab1/internal/service"
	"IntegrationLab1/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	service *service.Service
	log     *logger.Logger
	valid   *validator.Validate
}

func NewHandler(service *service.Service, log *logger.Logger, valid *validator.Validate) *Handler {
	return &Handler{
		service: service,
		log:     log,
		valid:   valid,
	}
}

func (h *Handler) submitCompletedDoc(c *gin.Context) {
	var input domain.UserRequest

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validateUserRequest(input); err != nil {

	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ответ": "данные получены успешно",
	})
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, message)
}

func (h *Handler) validateUserRequest(input domain.UserRequest) error {
	if err := h.valid.Struct(input); err != nil {
		return fmt.Errorf("Err(s):\n%+v\n", err)
	}
	return nil
}
