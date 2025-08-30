package routes

import (
	"github.com/gin-gonic/gin"

	httpNews "github.com/vinihss/bodego-api/internal/interfaces/http/news"
)

func RegisterNewsRoutes(r *gin.Engine, handler *httpNews.NewsHandler) {
	newsGroup := r.Group("/news")
	{
		newsGroup.POST("/", handler.Create)
		newsGroup.GET("/:id", handler.FindByID)
		newsGroup.GET("/", handler.GetAllNews)
		newsGroup.PUT("/:id", handler.Update)
		newsGroup.DELETE("/:id", handler.Delete)
	}
}