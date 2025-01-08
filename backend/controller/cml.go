package controller

import (
	"net/http"

	"github.com/Pleumzer/Dexon-Examination/entity"

	"github.com/gin-gonic/gin"
)

func CreateCml(c *gin.Context) {

	var cml entity.Cml

	if err := c.ShouldBindJSON(&cml); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&cml).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cml})

}
func GetCml(c *gin.Context) {
	var cml entity.Cml
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM cmls WHERE id = ?", id).Scan(&cml).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cml})
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
