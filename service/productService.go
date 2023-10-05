package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"myGo/models"
)

var DB *gorm.DB

func init() {
	dsn := "root:$Now2022@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				Colorful: true,
				LogLevel: logger.Info},
		),
	})
	if err != nil {
		log.Fatalf("fail to open mysql")
	}
	DB = db
}

func AddProduct(c *gin.Context) {
	data, _ := c.GetRawData()
	var pro models.Product
	_ = json.Unmarshal(data, &pro)
	result := DB.Create(&pro)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func UpdateProduct(c *gin.Context) {
	data, _ := c.GetRawData()
	var pro models.Product
	_ = json.Unmarshal(data, &pro)
	if pro.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	tx := DB.Model(&models.Product{}).Where("id = ?", pro.ID).UpdateColumns(pro)
	c.JSON(http.StatusOK, gin.H{"message": tx.RowsAffected})
}
func GetOne(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	var pro *models.Product
	DB.Where("id = ?", param).First(&pro)
	c.JSON(http.StatusOK, gin.H{"message": pro})
}
func GetAll(c *gin.Context) {
	var pros []*models.Product
	DB.Find(&pros)
	c.JSON(http.StatusOK, gin.H{"message": pros})
}
func DeleteProduct(c *gin.Context) {
	data, _ := c.GetRawData()
	var pro *models.Product
	json.Unmarshal(data, &pro)
	if pro.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	result := DB.Delete(&pro)
	c.JSON(http.StatusOK, gin.H{"message": result.RowsAffected})
}
