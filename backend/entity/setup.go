package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase () {
	database , err := gorm.Open(sqlite.Open("dababase.db") , &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

database.AutoMigrate(&Cml{})
database.AutoMigrate(&Info{})
database.AutoMigrate(&TEST_POINT{})
database.AutoMigrate(&THICKESS{})
db.AutoMigrate(&Info{}, &Cml{}, &TEST_POINT{}, &THICKESS{})
db =database
}