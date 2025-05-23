package category

import (
	models2 "github.com/Nurbol05/kolesa/shared/models"
)

type CategoryRepository interface {
	Create(category *models2.Category) error
	GetAll() ([]models2.Category, error)
	Update(id int, newName string) error
	Delete(id int) error
	GetCarsByCategoryID(categoryID int) ([]models2.Car, error)
}
type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(category *models2.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]models2.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Update(id int, newName string) error {
	return s.repo.Update(id, newName)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) GetCarsByCategoryID(categoryID int) ([]models2.Car, error) {
	return s.repo.GetCarsByCategoryID(categoryID)
}
