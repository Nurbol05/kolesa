package user

import (
	"errors"
	"gorm.io/gorm"
	"kolesa/user-service/models"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// CreateUser inserts a new user into the database
func (r UserRepositoryImpl) CreateUser(username, email, passwordHash string) error {
	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	return r.db.Create(&user).Error
}

// FindUserByEmail finds a user by email and returns the ID and hashed password
func (r UserRepositoryImpl) FindUserByEmail(email string) (int, string, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, "", errors.New("user not found")
		}
		return 0, "", err
	}
	return int(user.ID), user.PasswordHash, nil
}

// FindUserByID returns a user by ID
func (r UserRepositoryImpl) FindUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r UserRepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// UpdateUser updates an existing user's information in the database
func (r UserRepositoryImpl) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// DeleteUser deletes a user by ID
func (r UserRepositoryImpl) DeleteUser(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
