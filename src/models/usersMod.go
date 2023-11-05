package models

import "time"

type UserModel struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	PhoneNumber  string `gorm:"unique;not null;DEFAULT:0"`
	Email        string
	CreatedAt    time.Time `gorm:"default: current_timestamp"`
	UpdateAt     time.Time `gorm:"default: current_timestamp"`
	IsActive     bool      `gorm:"default: false"`
	IsPrime      bool      `gorm:"default: false"`
	Name         string
	PhotoProfile string
	Pin          string
	IdNumber     string
}
