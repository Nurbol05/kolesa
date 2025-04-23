package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtKey []byte

func init() {
	// Load JWT secret from environment variable
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

type UserRepository interface {
	CreateUser(username, email, passwordHash string) error
	FindUserByEmail(email string) (int, string, error)
	FindUserByID(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.CreateUser(username, email, string(hashedPassword))
}

func (s *UserService) Login(email, password string) (string, error) {
	userID, hashedPassword, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) UpdateUser(id int, username, email string) error {
	user, err := s.userRepo.FindUserByID(id)
	if err != nil {
		return err
	}

	user.Username = username
	user.Email = email

	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
