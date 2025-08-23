package http_interfaces_product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	controller *ProductController
}

func NewProductHandler(controller *ProductController) *ProductHandler {
	return &ProductHandler{controller: controller}
}

// Create handles the creation of a new product based on the provided JSON request. Returns the created product or an error.
// @Summary Create product
// @Description Creates a new product with the provided data
// @Tags Product
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product data"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} map[string]string
func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.CreateProduct(req)
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
// @Summary Delete product
// @Description Deletes a product by their ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product/{id} [delete]
// @Security BearerAuth
func (h *ProductHandler) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = h.controller.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// Update godoc
// @Summary Update product
// @Description Updates a product's information by their ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body UpdateProductRequest true "Updated product data"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product/{id} [put]
// @Security BearerAuth
func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.controller.UpdateProduct(uint(id), req)
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
// @Summary Get product by ID
// @Description Retrieves a product's information by their ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product/{id} [get]
// @Security BearerAuth
func (h *ProductHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := h.controller.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieves a paginated list of products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {array} ProductResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product [get]
// @Security BearerAuth
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := h.controller.GetAllProducts(page, size)
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
