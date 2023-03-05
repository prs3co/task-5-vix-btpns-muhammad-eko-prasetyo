package main

import (
	"github.com/prs3co/task-5-vix-btpns-muhammad-eko-prasetyo/models"
	"github.com/prs3co/task-5-vix-btpns-muhammad-eko-prasetyo/controllers/photocontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/photos", photocontroller.Index)
	r.GET("/api/photo/:id", photocontroller.Show)
	r.POST("/api/photo", photocontroller.Create)
	r.PUT("/api/photo/:id", photocontroller.Update)
	r.DELETE("/api/photo/:id", photocontroller.Delete)
	r.Run()
}