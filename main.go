package main

import (
	"Golang-for-Scalable-Backend/module/restaurant/transport/ginrestaurant"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL: ", db)

	// Lib Gin to write REST API
	r := gin.Default()
	v1 := r.Group("/api/v1")
	restaurants := v1.Group("/restaurants")

	// POST restaurants
	restaurants.POST("", ginrestaurant.CreateRestaurant(db))
	// DELETE RESTAURANTS
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
