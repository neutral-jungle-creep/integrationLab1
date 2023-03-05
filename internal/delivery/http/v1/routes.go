package v1

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	apiV1 := router.Group("/api/v1")
	{
		doc := apiV1.Group("/doc")
		{
			doc.POST("/get", h.submitCompletedDoc)
			doc.GET("/download/:filename", h.getDocFile)

		}
	}

	return router
}
