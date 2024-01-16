package main

import (
	"albumsApi/internal"
	"albumsApi/pkg/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.Default()

	internal.ConnectDatabase()

	/* routes */
	r.GET("/albums", handlers.FindAlbums)
	r.GET("/albums/:id", handlers.FindAlbum)
	r.POST("/albums", handlers.CreateAlbum)
	r.PATCH("/albums/:id", handlers.UpdateAlbum)
	r.DELETE("/albums/:id", handlers.DeleteAlbum)

	err := r.Run()

	if err != nil {
		panic("Failed to start server!")
	}
}
