package handlers

import (
	"albumsApi/internal"
	"albumsApi/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAlbumInput struct {
	Title    string `json:"title" binding:"required"`
	AuthorId string `json:"author_id" binding:"required"`
}

type UpdateAlbumInput struct {
	Title    string `json:"title"`
	AuthorId int    `json:"author_id"`
}

func FindAlbums(c *gin.Context) {
	var albums []models.Album
	internal.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func FindAlbum(c *gin.Context) {
	var album models.Album
	if err := internal.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found!"})
		return
	}

	fmt.Println(album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func CreateAlbum(c *gin.Context) {
	var input CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{Title: input.Title}
	internal.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	if err := internal.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found!"})
		return
	}

	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	internal.DB.Model(&album).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album
	if err := internal.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found!"})
		return
	}

	internal.DB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
