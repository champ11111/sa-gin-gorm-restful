package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/champ11111/sa-gin-gorm-restful/models" // Import your HouseKeeper model
	"github.com/champ11111/sa-gin-gorm-restful/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("PORT")

	fmt.Printf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// Migrate the HouseKeeper model to create the "housekeepers" table
	db.AutoMigrate(&models.HouseKeeper{})

	router := gin.New()
	api := router.Group("/api")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})

	routers.SetHouseKeeperRoutes(api, db)

	port := fmt.Sprintf(":%v", PORT)
	fmt.Println("Server Running on Port", port)
	http.ListenAndServe(port, router)
}
