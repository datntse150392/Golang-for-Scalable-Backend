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

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConn()

		var pagingData common.Page

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		pagingData.Fullfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
