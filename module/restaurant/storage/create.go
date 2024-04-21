package restaurantstorage

import (
	restaurantmodel "Golang-for-Scalable-Backend/module/restaurant/model"
	"context"
)

func (s *sqlStore) CreateRestaurant(context context.Context, data *restaurantmodel.Restaurant) error {
	return nil
}
