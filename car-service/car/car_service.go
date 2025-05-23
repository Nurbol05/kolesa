package car

import (
	"github.com/Nurbol05/kolesa/shared/models"
)

type CarRepository interface {
	Create(car *models.Car) error
	GetAll(params GetCarsParams) ([]models.Car, error) // жаңартылды
	GetByID(id int) (*models.Car, error)
	Update(car *models.Car) error
	Delete(id int) error
}

type GetCarsParams struct {
	Limit  int
	Page   int
	Filter string
}

type CarService struct {
	repo CarRepository
}

func NewCarService(repo CarRepository) *CarService {
	return &CarService{repo}
}

func (s *CarService) Create(car *models.Car) error {
	return s.repo.Create(car)
}

func (s *CarService) GetAll(params GetCarsParams) ([]models.Car, error) {
	return s.repo.GetAll(params)
}

func (s *CarService) GetByID(id int) (*models.Car, error) {
	return s.repo.GetByID(id)
}

func (s *CarService) Update(car *models.Car) error {
	return s.repo.Update(car)
}

func (s *CarService) Delete(id int) error {
	return s.repo.Delete(id)
}
