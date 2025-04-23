package category

import "kolesa/car"

type CategoryRepository interface {
	Create(category *Category) error
	GetAll() ([]Category, error)
	Update(id int, newName string) error
	Delete(id int) error
	GetCarsByCategoryID(categoryID int) ([]car.Car, error)
}
type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(category *Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Update(id int, newName string) error {
	return s.repo.Update(id, newName)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) GetCarsByCategoryID(categoryID int) ([]car.Car, error) {
	return s.repo.GetCarsByCategoryID(categoryID)
}
