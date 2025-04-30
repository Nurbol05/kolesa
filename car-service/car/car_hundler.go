package car

import (
	"github.com/gin-gonic/gin"
	"kolesa/car-service/models"
	"net/http"
	"strconv"
)

type CarHandler struct {
	service *CarService
}

func NewCarHandler(service *CarService) *CarHandler {
	return &CarHandler{service}
}

func (h *CarHandler) Create(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}
	c.JSON(http.StatusCreated, car)
}

func (h *CarHandler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter := c.Query("filter")

	params := GetCarsParams{
		Limit:  limit,
		Page:   page,
		Filter: filter,
	}

	cars, err := h.service.GetAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	car, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car.ID = id
	if err := h.service.Update(&car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}
	c.Status(http.StatusNoContent)
}
