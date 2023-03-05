package photocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prs3co/task-5-vix-btpns-muhammad-eko-prasetyo/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var photos []models.Photo

	models.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})

}

func Show(c *gin.Context) {
	var photo models.Photo
	id := c.Param("id")

	if err := models.DB.First(&photo, id).Error; err!= nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"photo": photo})

}

func Create(c *gin.Context) {
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
	}

	models.DB.Create(&photo)
	c.JSON(http.StatusOK, gin.H{"photo": photo})

}

func Update(c *gin.Context) {
	var photo models.Photo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
	}

	if models.DB.Model(&photo).Where("id = ?", id).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "update can't complete"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update is complete"})

}

func Delete(c *gin.Context) {
	var photo models.Photo

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&photo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "delete can't complete"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete is complete"})


}