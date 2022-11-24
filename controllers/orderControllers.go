package controllers

import (
	"LATIHAN1/database"
	"LATIHAN1/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)
func GetAllOrder(c *gin.Context) {
	var db = database.GetDB()

	Orders := []models.Order{}

	err := db.Find(&Orders).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Orders)
}

func GetOneOrder(c *gin.Context) {
	var db = database.GetDB()

	oneOrder := models.Order{}
	err := db.First(&oneOrder, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data one": oneOrder})
}

func CreateOrder(c *gin.Context) {
	var db = database.GetDB()
	var input = models.Order{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cariorder := models.Order{GormModel: input.GormModel, CustomerName: input.CustomerName, Items: input.Items}
	db.Create(&cariorder)
	
	c.JSON(http.StatusOK, cariorder)
}

func UpdateOrder(c *gin.Context) {
	var db = database.GetDB()

	order := models.Order{}
	err := db.First(&order, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&order).Updates(input)

	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{ Name: "id" }},
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}),		
	}).Create(&input)

	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{ Name: "id" }},
		DoUpdates: clause.AssignmentColumns([]string{"item_code", "description", "quantity"}),
	}).Create(&input.Items)
	c.JSON(http.StatusOK, input)
}

func DeleteOrder(c *gin.Context){
	var db = database.GetDB()
	orderdelete := models.Order{}
	err := db.First(&orderdelete, "Id = ?", c.Param("id")).Error;
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	db.Delete(&orderdelete)
	c.JSON(http.StatusOK, orderdelete)
}