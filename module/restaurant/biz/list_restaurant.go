package restaurantbiz

import (
	"Golang-for-Scalable-Backend/common"
	restaurantmodel "Golang-for-Scalable-Backend/module/restaurant/model"
	"context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Page,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Page,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
