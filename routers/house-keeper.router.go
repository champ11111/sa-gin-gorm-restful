package routers

import (
	"github.com/champ11111/sa-gin-gorm-restful/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetHouseKeeperRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	router.GET("houseKeepers", ctrls.GetHouseKeepers)
	router.GET("houseKeepers/:id", ctrls.GetHouseKeeper)
	router.POST("houseKeepers", ctrls.CreateHouseKeeper)
	router.PUT("houseKeepers/:id", ctrls.UpdateHouseKeeper)
	router.DELETE("houseKeepers/:id", ctrls.DeleteHouseKeeper)
}
