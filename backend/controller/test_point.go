package controller

import (
	"net/http"

	"github.com/Pleumzer/Dexon-Examination/entity"

	"github.com/gin-gonic/gin"
)

func CreateTestPoint(c *gin.Context) {

	var test_point entity.TEST_POINT

	if err := c.ShouldBindJSON(&test_point); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&test_point).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": test_point})

}
func GetTestPoint(c *gin.Context) {
	var test_point entity.TEST_POINT
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM test_points WHERE id = ?", id).Scan(&test_point).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test_point})
}

func ListTestPoint(c *gin.Context) {
	var test_points []entity.TEST_POINT
	if err := entity.DB().Raw("SELECT * FROM test_points").Scan(&test_points).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test_points})
}

func DeleteTestPoint(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM test_points WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "info not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateTestPoint(c *gin.Context) {

	var test_point entity.TEST_POINT

	if err := c.ShouldBindJSON(&test_point); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", test_point.ID).First(&test_point); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&test_point).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": test_point})

}
