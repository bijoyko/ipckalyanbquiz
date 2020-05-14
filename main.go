package main

import (
	"log"
	"os"

	"github.com/bijoyko/ipckalyanbquiz/controllers"
	"github.com/bijoyko/ipckalyanbquiz/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(gin.Logger())
	router.LoadHTMLGlob("view/*.html")
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/admin-bijoy", controllers.LoadAdmin)
	router.POST("/adminlogin", controllers.VerifyAdmin)
	router.POST("/updatequestions", controllers.UpdateQuestions)
	router.GET("/", controllers.MainPage)
	router.POST("/next", controllers.ValidateNames)
	router.POST("/submit", controllers.Form)
	router.POST("/viewTopScorers", controllers.ScoreView)
	router.GET("master-quiztable", controllers.Quiztable)
	router.Run(":" + port)
}
