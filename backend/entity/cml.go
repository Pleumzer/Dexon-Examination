package entity

import (
	"gorm.io/gorm"
)
type Cml struct {
	gorm.Model
	Cml_number              int    `json:"cml_number"`
	Cml_description         string `json:"cml_description"`
	Actual_Outside_Diameter int    `json:"actual_outside_diameter"`
	Design_thickness        int    `json:"design_thickness"`
	Structural_thickness    int    `json:"structural_thickness"`
	Required_thickness      int    `json:"required_thickness"`

	InfoID     uint         `json:"info_id"`
    TestPoints []TEST_POINT `gorm:"foreignKey:CmlID"`
}
