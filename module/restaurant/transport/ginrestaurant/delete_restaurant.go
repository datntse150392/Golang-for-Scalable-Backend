package ginrestaurant

import (
	"Golang-for-Scalable-Backend/common"
	"Golang-for-Scalable-Backend/component/appctx"
	restaurantbiz "Golang-for-Scalable-Backend/module/restaurant/biz"
	restaurantstorage "Golang-for-Scalable-Backend/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConn()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.Delete(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
