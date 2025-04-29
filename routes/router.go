package routes

import (
	"kolesa/car"
	"kolesa/category"
	"kolesa/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	var carRepo car.CarRepository = car.NewCarRepository(db)
	carService := car.NewCarService(carRepo)
	carHandler := car.NewCarHandler(carService)

	cars := r.Group("api/v1/cars")
	{
		cars.GET("/", carHandler.GetAll)       // Get all cars
		cars.GET("/:id", carHandler.GetByID)   // Get car by ID
		cars.POST("/", carHandler.Create)      // Create new car
		cars.PUT("/:id", carHandler.Update)    // Update car
		cars.DELETE("/:id", carHandler.Delete) // Delete car
	}

	var userRepo user.UserRepository = user.NewUserRepository(db)
	authService := user.NewUserService(userRepo)
	authHandler := user.NewUserHandler(authService)

	auth := r.Group("api/v1/user")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.GET("/", authHandler.GetAll)
		auth.PUT("/users/update", authHandler.UpdateUser)
		auth.DELETE("/users/delete/:id", authHandler.DeleteUser)
	}

	var categoryRepo category.CategoryRepository = category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := category.NewCategoryHandler(categoryService)

	cats := r.Group("api/v1/categories")
	{
		cats.GET("/", categoryHandler.GetAll)
		cats.POST("/", categoryHandler.Create)
		cats.PUT("/:id", categoryHandler.Update)
		cats.DELETE("/:id", categoryHandler.Delete)
		cats.GET("/:id/cars", categoryHandler.GetCarsByCategory)
	}
}
