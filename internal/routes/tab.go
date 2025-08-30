package routes

import (
	"github.com/gin-gonic/gin"

	httpTab "github.com/vinihss/bodego-api/internal/interfaces/http/tab"
)

func RegisterTabRoutes(r *gin.Engine, handler *httpTab.TabHandler) {
	tabGroup := r.Group("/tab")
	{
		tabGroup.POST("/", handler.OpenTab)                              // Open new tab
		tabGroup.GET("/", handler.GetAllTabs)                            // Get all tabs (paginated)
		tabGroup.GET("/:id", handler.FindByID)                           // Get tab by ID
		tabGroup.PUT("/:id", handler.UpdateTab)                          // Update tab description
		tabGroup.PUT("/:id/close", handler.CloseTab)                     // Close tab
		tabGroup.DELETE("/:id", handler.DeleteTab)                       // Delete tab
		tabGroup.GET("/user/:user_id", handler.GetTabsByUserID)          // Get tabs by user ID
		tabGroup.GET("/user/:user_id/open", handler.GetOpenTabsByUserID) // Get open tabs by user ID
	}
}
