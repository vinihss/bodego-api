package http_interfaces_drink

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinihss/bodego-api/internal/domain/drink"
)

type DrinkHandler struct {
	controller *DrinkController
}

func NewDrinkHandler(controller *DrinkController) *DrinkHandler {
	return &DrinkHandler{controller: controller}
}

func (h *DrinkHandler) Create(c *gin.Context) {
	var req CreateDrinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d, err := h.controller.createUC.Execute(drink.Drink{Name: req.Name, Price: req.Price})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, DrinkResponse{ID: d.ID, Name: d.Name, Price: d.Price})
}

func (h *DrinkHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	d, err := h.controller.findUC.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, DrinkResponse{ID: d.ID, Name: d.Name, Price: d.Price})
}

func (h *DrinkHandler) GetAll(c *gin.Context) {
	drinks, err := h.controller.findUC.ExecuteAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]DrinkResponse, len(drinks))
	for i, d := range drinks {
		resp[i] = DrinkResponse{ID: d.ID, Name: d.Name, Price: d.Price}
	}
	c.JSON(http.StatusOK, resp)
}

func (h *DrinkHandler) Update(c *gin.Context) {
	var req UpdateDrinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d, err := h.controller.updateUC.Execute(drink.Drink{ID: req.ID, Name: req.Name, Price: req.Price})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, DrinkResponse{ID: d.ID, Name: d.Name, Price: d.Price})
}

func (h *DrinkHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.controller.deleteUC.Execute(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
