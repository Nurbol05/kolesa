package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kolesa/models"
	"net/http"
)

type CategoryHandler struct {
	service *CategoryService
}

func NewCategoryHandler(service *CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	idParam := c.Param("id")

	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.Update(id, req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *CategoryHandler) GetCarsByCategory(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	cars, err := h.service.GetCarsByCategoryID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch cars by category"})
		return
	}

	c.JSON(http.StatusOK, cars)
}
