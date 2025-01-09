package controller

import (
	"net/http"

	"github.com/Pleumzer/Dexon-Examination/entity"

	"github.com/gin-gonic/gin"
)

func CreateCml(c *gin.Context) {
	var cml entity.Cml

	// ทำการ bind JSON เข้ากับโครงสร้าง Cml
	if err := c.ShouldBindJSON(&cml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่า InfoID ถูกส่งมาหรือไม่
	if cml.InfoID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "InfoID is required and cannot be zero"})
		return
	}

	// ดึงข้อมูล Info โดยใช้ InfoID ที่มาจาก Cml
	var info entity.Info
	if err := entity.DB().First(&info, cml.InfoID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Info not found for the provided InfoID"})
		return
	}

	// ตารางข้อมูล Pipe Size และ Actual Outside Diameter
	pipeData := map[float64]float64{
		0.125:  10.300,
		0.250:  13.700,
		0.357:  17.100,
		0.500:  21.300,
		0.750:  26.700,
		1.000:  33.400,
		1.250:  42.200,
		1.500:  48.300,
		2.000:  60.300,
		2.500:  73.000,
		3.000:  88.900,
		3.500:  101.600,
		4.000:  114.300,
		5.000:  141.300,
		6.000:  168.300,
		8.000:  219.100,
		10.000: 273.000,
		12.000: 323.800,
		14.000: 355.600,
		16.000: 406.400,
		18.000: 457.000,
	}

	// ตารางข้อมูล Structural Thickness ตาม Pipe Size
	structuralThickness := map[float64]float64{
		2.0:   1.80,  // <= 2
		3.0:   2.00,  // 3
		4.0:   2.30,  // 4
		6.0:   2.80,  // 6 <= pipe size <= 18
		20.0:  3.10,  // >= 20
	}

	// ตรวจสอบ Pipe Size และกำหนด Actual Outside Diameter
	if diameter, exists := pipeData[float64(info.PipeSize)]; exists {
		cml.Actual_Outside_Diameter = diameter // กำหนด Actual Outside Diameter
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pipe Size"})
		return
	}

	// กำหนด Structural Thickness ตาม Pipe Size
	switch {
	case info.PipeSize <= 2:
		cml.Structural_thickness = structuralThickness[2.0]
	case info.PipeSize == 3:
		cml.Structural_thickness = structuralThickness[3.0]
	case info.PipeSize == 4:
		cml.Structural_thickness = structuralThickness[4.0]
	case info.PipeSize >= 6 && info.PipeSize <= 18:
		cml.Structural_thickness = structuralThickness[6.0]
	case info.PipeSize >= 20:
		cml.Structural_thickness = structuralThickness[20.0]
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pipe Size for Structural Thickness"})
		return
	}

	// คำนวณ Design Thickness โดยใช้ฟังก์ชัน CalculateDesignThickness
	designThickness := CalculateDesignThickness(
		float64(info.Design_Pressure),       // designPressure
		cml.Actual_Outside_Diameter,        // actualOutsideDiameter
		float64(info.Stress),               // stress
		float64(info.Joint_Efficiency)/100, // jointEfficiency (แปลงจากเปอร์เซ็นต์เป็นทศนิยม)
	)

	// บันทึกค่าความหนาที่คำนวณได้ลงใน Cml
	cml.Design_thickness = designThickness

	if cml.Design_thickness > cml.Structural_thickness {
		cml.Required_thickness = cml.Design_thickness
	} else {
		cml.Required_thickness = cml.Structural_thickness
	}

	// บันทึกข้อมูล Cml ลงในฐานข้อมูล
	if err := entity.DB().Create(&cml).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cml})
}


func CalculateDesignThickness(designPressure, actualOutsideDiameter, stress, jointEfficiency float64) float64 {
	// คำนวณตัวส่วน
	denominator := (2 * stress * jointEfficiency) + (2 * designPressure * 0.4)
	// คำนวณตัวเศษและผลลัพธ์
	thickness := (designPressure * actualOutsideDiameter) / denominator
	return thickness
}



func GetCml(c *gin.Context) {
	var cmls []entity.Cml // เปลี่ยนจาก struct เป็น slice

	id := c.Param("id")
	// ใช้ Find เพื่อดึงข้อมูลทั้งหมดที่มี info_id ตรงกับ id
	if err := entity.DB().Where("info_id = ?", id).Find(&cmls).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cmls}) // คืนค่าข้อมูลทั้งหมด
}


func ListCml(c *gin.Context) {
	var cmls []entity.Cml
	if err := entity.DB().Raw("SELECT * FROM cmls").Scan(&cmls).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cmls})
}

func DeleteCml(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM cmls WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cml not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateCml(c *gin.Context) {

	var cml entity.Cml

	if err := c.ShouldBindJSON(&cml); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", cml.ID).First(&cml); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "cml not found"})

		return

	}

	if err := entity.DB().Save(&cml).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cml})

}


