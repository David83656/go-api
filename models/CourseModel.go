package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string
	Category    string
	Duration    int
	Description string
}
