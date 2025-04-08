package car

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var cars []Car
var carID = 1

func GetCars(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	limitStr := c.Query("limit")

	categoryID, _ := strconv.Atoi(categoryIDStr)
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 5
	}

	var filteredCars []Car
	for _, car := range cars {
		if categoryID > 0 && car.CategoryID != categoryID {
			continue
		}
		filteredCars = append(filteredCars, car)
		if len(filteredCars) >= limit {
			break
		}
	}

	c.JSON(http.StatusOK, filteredCars)
}

func GetCarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, car := range cars {
		if car.ID == id {
			c.JSON(http.StatusOK, car)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
}

func CreateCar(c *gin.Context) {
	var car Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if car.Brand == "" || car.Model == "" || car.Year == 0 || car.UserID == 0 || car.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	car.ID = carID
	carID++
	cars = append(cars, car)

	c.JSON(http.StatusCreated, car)
}

func UpdateCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате ID"})
		return
	}

	var updatedCar Car
	if err := c.ShouldBindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате JSON"})
		return
	}

	if updatedCar.Brand == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Машинаның бренді бос болмауы керек"})
		return
	}
	if updatedCar.Model == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Модель бос болмауы керек"})
		return
	}
	if updatedCar.Year <= 1900 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Жылы дұрыс емес"})
		return
	}
	if updatedCar.UserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID 0-ден үлкен болу керек"})
		return
	}
	if updatedCar.CategoryID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CategoryID 0-ден үлкен болу керек"})
		return
	}

	for i, car := range cars {
		if car.ID == id {
			updatedCar.ID = id
			cars[i] = updatedCar
			c.JSON(http.StatusOK, updatedCar)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Машина табылмады"})
}

func DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, car := range cars {
		if car.ID == id {
			cars = append(cars[:i], cars[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
}
