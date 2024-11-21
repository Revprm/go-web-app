package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Album represents an album model
type Album struct {
	id     int     `gorm:"primaryKey"`
	Title  string  `gorm:"size:128"`
	Artist string  `gorm:"size:255"`
	Price  float64 `gorm:"type:decimal(5,2)"`
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	db.AutoMigrate(&Album{})

	newAlbum := Album{Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 9.99}
	db.Create(&newAlbum)

	var albums []Album
	db.Find(&albums)
	for _, album := range albums {
		fmt.Printf("ID: %d, Title: %s, Artist: %s, Price: %.2f\n", album.id, album.Title, album.Artist, album.Price)
	}
}
