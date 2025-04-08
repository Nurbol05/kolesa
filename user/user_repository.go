package user

import (
	"errors"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// CreateUser inserts a new user into the database
func (r UserRepositoryImpl) CreateUser(username, email, passwordHash string) error {
	user := User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	return r.db.Create(&user).Error
}

// FindUserByEmail finds a user by email and returns the ID and hashed password
func (r UserRepositoryImpl) FindUserByEmail(email string) (int, string, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, "", errors.New("user not found")
		}
		return 0, "", err
	}
	return int(user.ID), user.PasswordHash, nil
}
