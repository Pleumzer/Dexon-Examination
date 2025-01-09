package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{}) // แก้ไขชื่อไฟล์จาก "dababase.db" เป็น "database.db"
	if err != nil {
		panic("failed to connect database: " + err.Error()) // แสดงข้อความผิดพลาดที่ชัดเจนขึ้น
	}

	// กำหนดค่า db หลังจากการเชื่อมต่อสำเร็จ
	db = database

	// ทำการ AutoMigrate หลังจาก db ถูกกำหนดค่า
	if err := db.AutoMigrate(&Cml{}, &Info{}, &TEST_POINT{}, &THICKESS{}); err != nil {
		panic("failed to migrate database: " + err.Error()) // แสดงข้อความผิดพลาดที่ชัดเจนขึ้น
	}
}
// func SetupDatabase () {
// 	database , err := gorm.Open(sqlite.Open("dababase.db") , &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// database.AutoMigrate(&Cml{})
// database.AutoMigrate(&Info{})
// database.AutoMigrate(&TEST_POINT{})
// database.AutoMigrate(&THICKESS{})
// db.AutoMigrate(&Info{}, &Cml{}, &TEST_POINT{}, &THICKESS{})
// db =database
// }