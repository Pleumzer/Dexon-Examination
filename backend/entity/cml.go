package entity

import (
	"gorm.io/gorm"
)
type Cml struct {
	gorm.Model
	Cml_number              int    `json:"cml_number"`
	Cml_description         string `json:"cml_description"`
	Actual_Outside_Diameter float64    `json:"actual_outside_diameter"`
	Design_thickness        float64    `json:"design_thickness"`
	Structural_thickness    float64    `json:"structural_thickness"`
	Required_thickness      float64    `json:"required_thickness"`

	InfoID     uint         `json:"info_id"`
    TestPoints []TEST_POINT `gorm:"foreignKey:CmlID"`
}
