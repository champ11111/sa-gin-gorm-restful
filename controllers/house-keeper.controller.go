package controllers

import (
	"fmt"
	"net/http"

	"github.com/champ11111/sa-gin-gorm-restful/models"
	"github.com/gin-gonic/gin"
)

// GET
func (db *DBController) GetHouseKeepers(c *gin.Context) {
	var houseKeepers []models.HouseKeeper
	if err := db.Database.Find(&houseKeepers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": houseKeepers})
}

// GET BY ID
func (db *DBController) GetHouseKeeper(c *gin.Context) {
	id := c.Param("id")
	var houseKeeper models.HouseKeeper

	db.Database.First(&houseKeeper, id)

	c.JSON(http.StatusOK, gin.H{"results": &houseKeeper})
}

// POST
func (db *DBController) CreateHouseKeeper(c *gin.Context) {
	var houseKeeper models.HouseKeeper
	c.BindJSON(&houseKeeper)

	if err := db.Database.Create(&houseKeeper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"results": &houseKeeper})
}

// PATCH
func (db *DBController) UpdateHouseKeeper(c *gin.Context) {
	id := c.Param("id")
	var houseKeeper models.HouseKeeper

	db.Database.First(&houseKeeper, id)

	fmt.Println(houseKeeper)
	c.BindJSON(&houseKeeper)

	if err := db.Database.Save(&houseKeeper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": &houseKeeper})
}

// DELETE
func (db *DBController) DeleteHouseKeeper(c *gin.Context) {

	id := c.Param("id")
	var houseKeeper models.HouseKeeper

	db.Database.Delete(&houseKeeper, id)

	c.JSON(http.StatusOK, gin.H{"results": "success"})
}
