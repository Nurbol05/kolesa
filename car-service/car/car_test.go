package car_test

import (
	car2 "github.com/Nurbol05/kolesa/car-service/car"
	"github.com/Nurbol05/kolesa/database"
	"github.com/Nurbol05/kolesa/shared/models"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB
var carRepo car2.CarRepository
var carService *car2.CarService

func setupTestDB(t *testing.T) {
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "postgres")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5433")
	_ = os.Setenv("DB_NAME", "kolesa_test")

	db, err := database.ConnectPostgres()
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	db.Migrator().DropTable(&models.Car{})
	db.AutoMigrate(&models.Car{})

	testDB = db
	carRepo = car2.NewCarRepository(testDB)
	carService = car2.NewCarService(carRepo)
}

func TestCreateCar(t *testing.T) {
	setupTestDB(t)

	newCar := models.Car{Brand: "BMW", Model: "X5", Year: 2022}
	err := carService.Create(&newCar)
	if err != nil {
		t.Fatalf("Create car failed: %v", err)
	}

	params := car2.GetCarsParams{
		Limit:  10,
		Page:   1,
		Filter: "",
	}
	cars, err := carService.GetAll(params)
	if err != nil {
		t.Fatalf("Failed to retrieve cars: %v", err)
	}

	if len(cars) != 1 {
		t.Fatalf("Expected 1 car in the database, found %d", len(cars))
	}
}

func TestUpdateCar(t *testing.T) {
	setupTestDB(t)

	newCar := models.Car{Brand: "BMW", Model: "X5", Year: 2022}
	_ = carService.Create(&newCar)

	params := car2.GetCarsParams{
		Limit:  10,
		Page:   1,
		Filter: "",
	}
	cars, _ := carService.GetAll(params)
	if len(cars) != 1 {
		t.Fatal("Expected one car in DB")
	}

	carToUpdate := cars[0]
	carToUpdate.Brand = "Mercedes"
	carToUpdate.Model = "GLA"
	err := carService.Update(&carToUpdate)
	if err != nil {
		t.Fatalf("Update car failed: %v", err)
	}

	updatedCar, _ := carRepo.GetByID(carToUpdate.ID)
	if updatedCar.Brand != "Mercedes" || updatedCar.Model != "GLA" {
		t.Fatal("Car not updated correctly")
	}
}

func TestDeleteCar(t *testing.T) {
	setupTestDB(t)

	newCar := models.Car{Brand: "BMW", Model: "X5", Year: 2022}
	_ = carService.Create(&newCar)

	params := car2.GetCarsParams{
		Limit:  10,
		Page:   1,
		Filter: "",
	}
	cars, _ := carService.GetAll(params)
	carID := cars[0].ID

	err := carService.Delete(carID)
	if err != nil {
		t.Fatalf("Delete car failed: %v", err)
	}

	_, err = carRepo.GetByID(carID)
	if err == nil {
		t.Fatal("Expected error when retrieving deleted car")
	}
}
