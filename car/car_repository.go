package car

import (
	"gorm.io/gorm"
)

type CarRepositoryImpl struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepositoryImpl {
	return &CarRepositoryImpl{db}
}

// Используем структуру Car из car.go
func (r CarRepositoryImpl) Create(car *Car) error {
	return r.db.Create(car).Error
}

func (r CarRepositoryImpl) GetAll(params GetCarsParams) ([]Car, error) {
	var cars []Car
	query := r.db.Model(&Car{})

	// Фильтр қолдану (мысалы, бренд)
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

func (r CarRepositoryImpl) GetByID(id int) (*Car, error) {
	var car Car
	err := r.db.First(&car, id).Error
	return &car, err
}

func (r CarRepositoryImpl) Update(car *Car) error {
	return r.db.Save(car).Error
}

func (r CarRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&Car{}, id).Error
}
