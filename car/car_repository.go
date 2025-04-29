package car

import (
	"gorm.io/gorm"
	"kolesa/models"
)

type CarRepositoryImpl struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepositoryImpl {
	return &CarRepositoryImpl{db}
}

func (r CarRepositoryImpl) Create(car *models.Car) error {
	return r.db.Create(car).Error
}

func (r CarRepositoryImpl) GetAll(params GetCarsParams) ([]models.Car, error) {
	var cars []models.Car
	query := r.db.Model(&models.Car{})

	if params.Filter != "" {
		query = query.Where("brand ILIKE ?", "%"+params.Filter+"%")
	}

	// Пагинация
	offset := (params.Page - 1) * params.Limit
	err := query.Limit(params.Limit).Offset(offset).Find(&cars).Error

	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (r CarRepositoryImpl) GetByID(id int) (*models.Car, error) {
	var car models.Car
	err := r.db.First(&car, id).Error
	return &car, err
}

func (r CarRepositoryImpl) Update(car *models.Car) error {
	return r.db.Save(car).Error
}

func (r CarRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.Car{}, id).Error
}
