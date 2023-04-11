package get

import (
	"main/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetHumi(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var Temp []entities.Temp

		err := ctx.ShouldBindJSON(&Temp)
		farmlandname := ctx.Param("farmlandname")

		if err != nil {
			db.Debug().Table("farmlands").
				Select([]string{"temphums.temp", "temphums.humi"}).
				Joins("JOIN temphums ON farmlands.id = temphums.farmland_id").
				Where("farmlands.farmlandname = ?", farmlandname).
				Scan(&Temp)
			ctx.JSON(http.StatusOK, Temp)
		}
	}

}