package entity

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "m"
	Female Gender = "f"
	Other  Gender = "o"
)

func (g *Gender) Scan(value interface{}) error {
	*g = Gender(value.([]byte))
	return nil
}

func (g Gender) Value() (driver.Value, error) {
	return string(g), nil
}

type Customer struct {
	gorm.Model
	Fname     string
	Lname     string
	Gender    Gender `json:"gender" gorm:"type:enum('m', 'f', 'o');"`
	BirthDate time.Time
}
