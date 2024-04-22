package restaurantbiz

import (
	restaurantmodel "Golang-for-Scalable-Backend/module/restaurant/model"
	"context"
	"errors"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) Delete(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data has been deleted")
	}

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
