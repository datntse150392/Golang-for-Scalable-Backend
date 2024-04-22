package ginrestaurant

import (
	"Golang-for-Scalable-Backend/common"
	"Golang-for-Scalable-Backend/component/appctx"
	restaurantbiz "Golang-for-Scalable-Backend/module/restaurant/biz"
	restaurantmodel "Golang-for-Scalable-Backend/module/restaurant/model"
	restaurantstorage "Golang-for-Scalable-Backend/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConn()
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
