package v1

import (
	"IntegrationLab1/internal/domain"
	"IntegrationLab1/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (h *Handler) SubmitCompletedDoc(c *gin.Context) {
	var input domain.UserRequest

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: make separate func for validate
	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ответ": "данные получены успешно",
	})
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, message)
}
