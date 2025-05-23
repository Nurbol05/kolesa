package category_test

import (
	"github.com/Nurbol05/kolesa/category"
	"github.com/Nurbol05/kolesa/database"
	"github.com/Nurbol05/kolesa/shared/models"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB
var categoryRepo category.CategoryRepositoryImpl
var categoryService *category.CategoryService

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

	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})

	testDB = db
	categoryRepo = *category.NewCategoryRepository(testDB)
	categoryService = category.NewCategoryService(categoryRepo)
}

func TestCreateCategory(t *testing.T) {
	setupTestDB(t)

	category := models.Category{Name: "SUV"}
	err := categoryRepo.Create(&category)
	if err != nil {
		t.Fatalf("Create category failed: %v", err)
	}

	categories, _ := categoryRepo.GetAll()
	if len(categories) != 1 {
		t.Fatalf("Expected 1 category in the database, found %d", len(categories))
	}
}

func TestUpdateCategory(t *testing.T) {
	setupTestDB(t)

	category := models.Category{Name: "SUV"}
	_ = categoryRepo.Create(&category)

	categories, _ := categoryRepo.GetAll()
	if len(categories) != 1 {
		t.Fatal("Expected one category in DB")
	}

	categoryToUpdate := categories[0]
	categoryToUpdate.Name = "Luxury SUV"
	err := categoryRepo.Update(categoryToUpdate.ID, categoryToUpdate.Name)
	if err != nil {
		t.Fatalf("Update category failed: %v", err)
	}

	updatedCategory, _ := categoryRepo.GetAll()
	if updatedCategory[0].Name != "Luxury SUV" {
		t.Fatal("Category not updated correctly")
	}
}

func TestDeleteCategory(t *testing.T) {
	setupTestDB(t)

	category := models.Category{Name: "SUV"}
	_ = categoryRepo.Create(&category)

	categories, _ := categoryRepo.GetAll()
	categoryID := categories[0].ID

	err := categoryRepo.Delete(categoryID)
	if err != nil {
		t.Fatalf("Delete category failed: %v", err)
	}

	cars, err := categoryRepo.GetCarsByCategoryID(categoryID)
	if err != nil {
		t.Fatalf("Unexpected error when retrieving cars: %v", err)
	}
	if len(cars) != 0 {
		t.Fatalf("Expected 0 cars for deleted category, but got %d", len(cars))
	}
}
