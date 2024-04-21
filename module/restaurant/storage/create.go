package restaurantstorage

import (
	restaurantmodel "Golang-for-Scalable-Backend/module/restaurant/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
