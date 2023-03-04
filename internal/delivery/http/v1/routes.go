package v1

import "github.com/gin-gonic/gin"

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/", "./frontend")

	doc := router.Group("/api/v1")
	{
		doc.POST("/get-doc", h.submitCompletedDoc)
	}

	return router
}
