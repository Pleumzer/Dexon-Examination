package entity

import "time"

import (
	"gorm.io/gorm"
)

type THICKESS struct {
	gorm.Model
	Inspection_date  time.Time `json:"inspection_date"`
	Actual_thickness int       `json:"actual_thickness"`

	TestPointID uint `json:"test_point_id"`
}
