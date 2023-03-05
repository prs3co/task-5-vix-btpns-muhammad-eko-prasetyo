package photocontroller

import (
	"net/http"
	"github.com/prs3co/task-5-vix-btpns-muhammad-eko-prasetyo/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var photos []models.Photo

	models.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})

}
func Show(c *gin.Context) {

}
func Create(c *gin.Context) {

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}