package models

type Car struct {
	ID         int    `json:"id"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
}
