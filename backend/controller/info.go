package controller

import (
	"net/http"

	"github.com/Pleumzer/Dexon-Examination/entity"

	"github.com/gin-gonic/gin"
)

func CreateInfo(c *gin.Context) {

	var info entity.Info

	if err := c.ShouldBindJSON(&info); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&info).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": info})

}
func GetInfo(c *gin.Context) {
	var info entity.Info
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM infos WHERE id = ?", id).Scan(&info).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info})
}

func ListInfo(c *gin.Context) {
	var infos []entity.Info
	if err := entity.DB().Raw("SELECT * FROM infos").Scan(&infos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": infos})
}

func DeleteInfo(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM infos WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "info not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateInfo(c *gin.Context) {

	var info entity.Info

	if err := c.ShouldBindJSON(&info); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", info.ID).First(&info); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&info).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": info})

}
