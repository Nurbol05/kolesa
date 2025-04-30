package category

import (
	"gorm.io/gorm"
	models2 "kolesa/car-service/models"
	"kolesa/models"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db}
}

func (r CategoryRepositoryImpl) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r CategoryRepositoryImpl) Update(id int, newName string) error {
	return r.db.Model(&models.Category{}).Where("id = ?", id).Update("name", newName).Error
}

func (r CategoryRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.Category{}, id).Error
}

func (r CategoryRepositoryImpl) GetCarsByCategoryID(categoryID int) ([]models2.Car, error) {
	var cars []models2.Car
	err := r.db.Where("category_id = ?", categoryID).Find(&cars).Error
	return cars, err
}
