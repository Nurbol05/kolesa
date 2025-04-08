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

func (r CarRepositoryImpl) GetAll() ([]Car, error) {
	var cars []Car
	err := r.db.Find(&cars).Error
	return cars, err
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
