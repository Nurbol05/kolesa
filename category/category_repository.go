package category

import (
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db}
}

func (r CategoryRepositoryImpl) Create(category *Category) error {
	return r.db.Create(category).Error
}

func (r CategoryRepositoryImpl) GetAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r CategoryRepositoryImpl) Update(id int, newName string) error {
	return r.db.Model(&Category{}).Where("id = ?", id).Update("name", newName).Error
}

func (r CategoryRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&Category{}, id).Error
}
