package category

import (
	"kolesa/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	Update(id int, newName string) error
	Delete(id int) error
	GetCarsByCategoryID(categoryID int) ([]models.Car, error)
}
type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Update(id int, newName string) error {
	return s.repo.Update(id, newName)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) GetCarsByCategoryID(categoryID int) ([]models.Car, error) {
	return s.repo.GetCarsByCategoryID(categoryID)
}
