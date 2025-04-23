package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"kolesa/car"
	"kolesa/category"
	"kolesa/database"
	"kolesa/routes"
	"kolesa/user"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Орта айнымалыларын орнату
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "kolesa")
	os.Setenv("DB_PORT", "5432")

	os.Exit(m.Run())
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	database.ConnectPostgres()
	routes.SetupRoutes(router, database.DB)
	return router
}

func initTestDB() {
	database.ConnectPostgres()
	database.DB.Exec("DELETE FROM users")
	database.DB.Exec("DELETE FROM categories")
	database.DB.Exec("DELETE FROM cars")
}

// --- USER TESTS ---
func TestRegisterUser(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	u := user.User{Username: "tester", Email: "test@example.com", PasswordHash: "password123"}
	body, _ := json.Marshal(u)

	req, _ := http.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// --- CATEGORY TESTS ---
func TestCreateCategory(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	c := category.Category{Name: "Luxury"}
	body, _ := json.Marshal(c)

	req, _ := http.NewRequest("POST", "/api/v1/categories/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// --- CAR TESTS ---
func TestCreateCar(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	// 1. Пайдаланушы тіркеу
	user := user.User{Username: "testuser", Email: "testuser@example.com", PasswordHash: "password123"}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Тіркелген пайдаланушының ID-сін алу
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 but got %d", w.Code)
	}

	var userResponse map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&userResponse); err != nil {
		t.Fatalf("Error decoding user response: %v", err)
	}

	// 'id' кілтінің болуын тексереміз
	userID, ok := userResponse["id"].(float64)
	if !ok {
		t.Fatalf("Expected 'id' to be present in user response, got: %v", userResponse["id"])
	}

	// 2. Категория жасау
	category := category.Category{Name: "Luxury"}
	body, _ = json.Marshal(category)
	req, _ = http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Категорияның ID-сін алу
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 but got %d", w.Code)
	}

	var categoryResponse map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&categoryResponse); err != nil {
		t.Fatalf("Error decoding category response: %v", err)
	}

	// 'id' кілтінің болуын тексереміз
	categoryID, ok := categoryResponse["id"].(float64)
	if !ok {
		t.Fatalf("Expected 'id' to be present in category response, got: %v", categoryResponse["id"])
	}

	// 3. Машина жасау
	car := car.Car{Brand: "Toyota", Model: "Camry", Year: 2023, UserID: int(userID), CategoryID: int(categoryID)}
	body, _ = json.Marshal(car)
	req, _ = http.NewRequest("POST", "/api/v1/cars", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Жауапты тексеру
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 but got %d", w.Code)
	}
	assert.Equal(t, http.StatusCreated, w.Code)
}
