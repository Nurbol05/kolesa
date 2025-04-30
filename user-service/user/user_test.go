package user_test

import (
	"kolesa/database"
	"kolesa/user-service/models"
	user2 "kolesa/user-service/user"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB
var userRepo user2.UserRepository
var userService *user2.UserService

func setupTestDB(t *testing.T) {
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "postgres")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_NAME", "kolesa_test")

	db, err := database.ConnectPostgres()
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	testDB = db
	userRepo = user2.NewUserRepository(testDB)
	userService = user2.NewUserService(userRepo)
}

func TestRegisterAndLogin(t *testing.T) {
	setupTestDB(t)

	err := userService.Register("testuser", "test@example.com", "password123")
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	token, err := userService.Login("test@example.com", "password123")
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if token == "" {
		t.Fatal("Expected JWT token, got empty string")
	}
}

func TestUpdateUser(t *testing.T) {
	setupTestDB(t)
	_ = userService.Register("user2", "user2@example.com", "pass")

	users, _ := userService.GetAll()
	if len(users) != 1 {
		t.Fatal("Expected one user in DB")
	}

	err := userService.UpdateUser(int(users[0].ID), "updatedUser", "updated@example.com")
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}

	updatedUser, _ := userRepo.FindUserByID(int(users[0].ID))
	if updatedUser.Username != "updatedUser" || updatedUser.Email != "updated@example.com" {
		t.Fatal("User not updated correctly")
	}
}

func TestDeleteUser(t *testing.T) {
	setupTestDB(t)
	_ = userService.Register("todelete", "delete@example.com", "pass")

	users, _ := userService.GetAll()
	userID := int(users[0].ID)

	err := userService.DeleteUser(userID)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

	_, err = userRepo.FindUserByID(userID)
	if err == nil {
		t.Fatal("Expected error when retrieving deleted user")
	}
}
