package models

import "gorm.io/gorm"

type File struct {
	ID            int
	Name          string
	Path          string
	PublicacionID int
	gorm.Model    `json:"-"`
}

type User struct {
	ID    int
	email string
	gorm.Model
}
