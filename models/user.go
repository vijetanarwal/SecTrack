package models

import "time"

type User struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    Email        string    `json:"email" gorm:"unique;not null"`
    PasswordHash string    `json:"-" gorm:"not null"`
    Role         string    `json:"role" gorm:"default:'viewer'"`
    CreatedAt    time.Time `json:"created_at"`
}