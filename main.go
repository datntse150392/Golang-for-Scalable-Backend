package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

/*
Vì sao ở đây lại phải tạo 1 struct riêng cho Update
-> Đầu tiên: là vì một dữ liệu được update ở Go nếu có trường hợp này sẽ bỏ qua: String bị "", boolean: flase, number: 0 thì sẽ không update
-> Vì vậy chỉ cần trỏ một con trỏ vào để cho nó có data, chỉ khi không có trỏ tới nào thì data mới không được update
*/
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

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
	restaurants.POST("", func(c *gin.Context) {
		var data Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data": &data,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// GET RESTAURANT BY ID
	restaurants.GET("/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data Restaurant

		db.Where("id = ?", id).Find(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// REST ALL RESTAURANT
	restaurants.GET("", func(c *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var pageData Paging

		if pageData.Page <= 0 {
			pageData.Page = 1
		}
		if pageData.Limit <= 0 {
			pageData.Limit = 5
		}

		var data []Restaurant

		db.Order("id asc").Offset((pageData.Page - 1) * pageData.Limit).Limit(pageData.Limit).Find(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// Update info Restaurant
	restaurants.PUT("/:id", func(c *gin.Context) {
		var data RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		db.Where("id = ?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// DELETE RESTAURANT BY ID
	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"data": id,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
