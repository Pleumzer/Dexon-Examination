package entity

import (
	"gorm.io/gorm"
)

type TEST_POINT struct {
	gorm.Model
	Tp_number      int    `json:"tp_number"`
	Tp_description int    `json:"tp_description"`
	Note           string `json:"note"`

	CmlID       uint        `json:"cml_id"`
    Thicknesses []THICKESS  `gorm:"foreignKey:TestPointID"`
}