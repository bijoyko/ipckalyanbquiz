package main

import (
	"log"
	"os"

	"github.com/bijoyko/ipckalyanbquiz/controllers"
	"github.com/bijoyko/ipckalyanbquiz/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("*.html")
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/", controllers.MainPage)
	router.POST("/next", controllers.ValidateNames)
	router.POST("/submit", controllers.Form)
	router.Run(":" + port)
}
