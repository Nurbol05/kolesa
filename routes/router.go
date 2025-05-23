package routes

import (
	"github.com/Nurbol05/kolesa/category"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	//var carRepo car2.CarRepository = car2.NewCarRepository(db)
	//carService := car2.NewCarService(carRepo)
	//carHandler := car2.NewCarHandler(carService)
	//
	//cars := r.Group("api/v1/cars")
	//{
	//	cars.GET("/", carHandler.GetAll)
	//	cars.GET("/:id", carHandler.GetByID)
	//	cars.POST("/", carHandler.Create)
	//	cars.PUT("/:id", carHandler.Update)
	//	cars.DELETE("/:id", carHandler.Delete)
	//}

	//var userRepo user2.UserRepository = user2.NewUserRepository(db)
	//authService := user2.NewUserService(userRepo)
	//authHandler := user2.NewUserHandler(authService)
	//
	//auth := r.Group("api/v1/user")
	//{
	//	auth.POST("/register", authHandler.Register)
	//	auth.POST("/login", authHandler.Login)
	//	auth.GET("/", authHandler.GetAll)
	//	auth.PUT("/users/update", authHandler.UpdateUser)
	//	auth.DELETE("/users/delete/:id", authHandler.DeleteUser)
	//}

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
