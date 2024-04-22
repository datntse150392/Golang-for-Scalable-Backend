package restaurantmodel

import (
	"Golang-for-Scalable-Backend/common"
	"errors"
	"strings"
)

type Restaurant struct {
	common.SQLModel
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

type RestaurantCreate struct {
	common.SQLModel
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

var (
	ErrNameIsEmpty = errors.New("Name is cannot be empty")
)
