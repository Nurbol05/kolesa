package user

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}
