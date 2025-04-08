package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var categories []Category
var categoryID = 1

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = categoryID
	categoryID++
	categories = append(categories, category)

	c.JSON(http.StatusCreated, category)
}
