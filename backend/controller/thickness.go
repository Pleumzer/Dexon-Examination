package controller

import (
	"net/http"

	"github.com/Pleumzer/Dexon-Examination/entity"

	"github.com/gin-gonic/gin"
)

func CreateThickness(c *gin.Context) {

	var thickness entity.THICKESS

	if err := c.ShouldBindJSON(&thickness); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&thickness).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": thickness})

}
func GetThickness(c *gin.Context) {
	var thickness entity.THICKESS
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM thicknesses WHERE id = ?", id).Scan(&thickness).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": thickness})
}

func ListThickness(c *gin.Context) {
	var thicknesses []entity.THICKESS
	if err := entity.DB().Raw("SELECT * FROM thicknesses").Scan(&thicknesses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": thicknesses})
}

func DeleteThickness(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM thicknesses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thickness not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateThickness(c *gin.Context) {

	var thickness entity.THICKESS

	if err := c.ShouldBindJSON(&thickness); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", thickness.ID).First(&thickness); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "thickness not found"})

		return

	}

	if err := entity.DB().Save(&thickness).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": thickness})

}
