package routes

import (
	"github.com/gin-gonic/gin"
	http_drink "github.com/vinihss/bodego-api/internal/interfaces/http/drink"
)

func RegisterDrinkRoutes(r *gin.Engine, handler *http_drink.DrinkHandler) {
	drinks := r.Group("/drinks")
	{
		drinks.POST("/", handler.Create)
		drinks.GET("/", handler.GetAll)
		drinks.GET(":id", handler.GetByID)
		drinks.PUT("/", handler.Update)
		drinks.DELETE(":id", handler.Delete)
	}
}
