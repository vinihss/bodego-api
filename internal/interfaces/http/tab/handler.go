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

// OpenTab godoc
// @Summary Open a new tab
// @Description Opens a new tab for a user
// @Tags Tab
// @Accept json
// @Produce json
// @Param tab body OpenTabRequest true "Tab data"
// @Success 201 {object} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab [post]
// @Security BearerAuth
func (h *TabHandler) OpenTab(c *gin.Context) {
	var req OpenTabRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.OpenTab(req)
	if err != nil {
		if strings.Contains(err.Error(), "already has an open tab") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// CloseTab godoc
// @Summary Close a tab
// @Description Closes an open tab by its ID
// @Tags Tab
// @Accept json
// @Produce json
// @Param id path int true "Tab ID"
// @Success 200 {object} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/{id}/close [put]
// @Security BearerAuth
func (h *TabHandler) CloseTab(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := h.controller.CloseTab(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "already closed") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "tab not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateTab godoc
// @Summary Update tab
// @Description Updates a tab's description by its ID
// @Tags Tab
// @Accept json
// @Produce json
// @Param id path int true "Tab ID"
// @Param tab body UpdateTabRequest true "Updated tab data"
// @Success 200 {object} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/{id} [put]
// @Security BearerAuth
func (h *TabHandler) UpdateTab(c *gin.Context) {
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
		if strings.Contains(err.Error(), "cannot update a closed tab") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "tab not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// FindByID godoc
// @Summary Get tab by ID
// @Description Retrieves a tab's information by its ID
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
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "tab not found"})
			return
		}
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

// GetTabsByUserID godoc
// @Summary Get tabs by user ID
// @Description Retrieves all tabs for a specific user
// @Tags Tab
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/user/{user_id} [get]
// @Security BearerAuth
func (h *TabHandler) GetTabsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	res, err := h.controller.GetTabsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetOpenTabsByUserID godoc
// @Summary Get open tabs by user ID
// @Description Retrieves all open tabs for a specific user
// @Tags Tab
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} TabResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/user/{user_id}/open [get]
// @Security BearerAuth
func (h *TabHandler) GetOpenTabsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	res, err := h.controller.GetOpenTabsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteTab godoc
// @Summary Delete tab
// @Description Deletes a tab by its ID
// @Tags Tab
// @Accept json
// @Produce json
// @Param id path int true "Tab ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tab/{id} [delete]
// @Security BearerAuth
func (h *TabHandler) DeleteTab(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.controller.DeleteTab(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "tab not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tab deleted successfully"})
}
