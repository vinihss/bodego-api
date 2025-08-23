package http_interfaces_tab

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type TabHandler struct {
	controller *TabController
}

func NewTabHandler(controller *TabController) *TabHandler {
	return &TabHandler{controller: controller}
}

// Create godoc
// @Summary Add tab product
// @Description Creates a favorite for a given tab and product
// @Tags Tab
// @Accept json
// @Produce json
// @Param favorite body CreateTabRequest true "Tab data"
// @Success 200 {object} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab [post]
// @Security BearerAuth
func (h *TabHandler) Create(c *gin.Context) {
	var req CreateTabRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.CreateTab(req)
	if err != nil {
		if isUniqueEmailErr(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete tab
// @Description Deletes a tab by their ID
// @Tags Tab
// @Accept json
// @Produce json
// @Param id path int true "Tab ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/{id} [delete]
// @Security BearerAuth
func (h *TabHandler) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = h.controller.DeleteTab(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tab deleted successfully"})
}

// Update godoc
// @Summary Update tab
// @Description Updates a tab's information by their ID
// @Tags Tab
// @Accept json
// @Produce json
// @Param id path int true "Tab ID"
// @Param tab body UpdateTabRequest true "Updated tab data"
// @Success 200 {object} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/{id} [put]
// @Security BearerAuth
func (h *TabHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateTabRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.UpdateTab(uint(id), req)
	if err != nil {
		if isUniqueEmailErr(err) {
			c.JSON(http.StatusConflict, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// FindByID godoc
// @Summary Get tab by ID
// @Description Retrieves a tab's information by their ID
// @Tags Tab
// @Accept json
// @Produce json
// @Param id path int true "Tab ID"
// @Success 200 {object} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/{id} [get]
// @Security BearerAuth
func (h *TabHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := h.controller.GetTab(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllTabs godoc
// @Summary Get all tabs
// @Description Retrieves a paginated list of tabs
// @Tags Tab
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {array} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab [get]
// @Security BearerAuth
func (h *TabHandler) GetAllTabs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := h.controller.GetAllTabs(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func isUniqueEmailErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())

	return (strings.Contains(msg, "unique") || strings.Contains(msg, "duplicate") || strings.Contains(msg, "already exists")) &&
		strings.Contains(msg, "email")
}
