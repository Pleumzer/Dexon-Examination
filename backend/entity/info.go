package entity

import "time"

import (
	"gorm.io/gorm"
)

// Info represents the entity structure for the given table.
type Info struct {
	gorm.Model
	Line_Number           string    `json:"line_number"`
	Location              string    `json:"location"`
	From                  string    `json:"from"`
	To                    string    `json:"to"`
	DrawingNumber         string    `json:"drawing_number"`
	Service               string    `json:"service"`
	Material              string    `json:"material"`
	InService_Date        time.Time `json:"inservice_date"`
	PipeSize              int       `json:"pipe_size"`
	OriginalThickness     int       `json:"original_thickness"`
	Stress                int       `json:"stress"`
	Joint_Efficiency      int       `json:"joint_efficiency"`
	CA                    int       `json:"ca"`
	Design_Life           int       `json:"design_life"`
	Design_Pressure       int       `json:"design_pressure"`
	Operating_Pressure    int       `json:"operating_pressure"`
	Design_Temperature    int       `json:"design_temperature"`
	Operating_Temperature int       `json:"operating_temperature"`

	Cmls []Cml `gorm:"foreignKey:InfoID"`
}



