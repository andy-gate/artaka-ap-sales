package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andy-gate/artaka-ap-sales/controllers"
	"github.com/andy-gate/artaka-ap-sales/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())

	if err := godotenv.Load(`.env`); err != nil {
		panic(err)
	}

	models.InitGormPostgres()
	defer models.MPosGORM.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")
	api.POST("/send_sales", controllers.SendSales)
	
	fmt.Printf("Listening to port %s", os.Getenv("PORT1"))
	router.Run(":" + os.Getenv("PORT1"))
}