package http_interfaces_news

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NewsHandler struct {
	controller *NewsController
}

func NewNewsHandler(controller *NewsController) *NewsHandler {
	return &NewsHandler{controller: controller}
}

// Create godoc
// @Summary Add news item
// @Description Creates a news item
// @Tags News
// @Accept json
// @Produce json
// @Param news body CreateNewsRequest true "News data"
// @Success 200 {object} NewsResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news [post]
// @Security BearerAuth
func (h *NewsHandler) Create(c *gin.Context) {
	var req CreateNewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.CreateNews(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete news item
// @Description Deletes a news item by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news/{id} [delete]
// @Security BearerAuth
func (h *NewsHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = h.controller.DeleteNews(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}

// Update godoc
// @Summary Update news item
// @Description Updates a news item by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Param news body UpdateNewsRequest true "Updated news data"
// @Success 200 {object} NewsResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news/{id} [put]
// @Security BearerAuth
func (h *NewsHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateNewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.UpdateNews(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// FindByID godoc
// @Summary Get news item by ID
// @Description Retrieves a news item by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} NewsResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news/{id} [get]
// @Security BearerAuth
func (h *NewsHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := h.controller.GetNews(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllNews godoc
// @Summary Get all news items
// @Description Retrieves a paginated list of news items
// @Tags News
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {array} NewsResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news [get]
// @Security BearerAuth
func (h *NewsHandler) GetAllNews(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := h.controller.GetAllNews(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}