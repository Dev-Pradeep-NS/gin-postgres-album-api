package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type album struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *gorm.DB
var err error

func main() {
	dsn := "host=localhost user=albumuser password=password dbname=albumdb port=5432 sslmode=disable TimeZone=Asia/Calcutta"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&album{})

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8000")
}

func getAlbums(c *gin.Context) {
	var albums []album
	db.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var album album
	result := db.First(&album, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func updateAlbum(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var album album
	result := db.First(&album, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	db.Model(&album).Updates(updatedAlbum)
	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbum(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var album album
    if err := db.First(&album, id).Error; err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
        return
    }

    if err := db.Delete(&album).Error; err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
