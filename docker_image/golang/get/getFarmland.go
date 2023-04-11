package get

import (
	"main/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFarmland(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var farmaland []entities.Farmland

		err := ctx.ShouldBindJSON(&farmaland)
		username := ctx.Param("username")

		if err != nil {
			db.Debug().
			Table("agrists").
			Select([]string{ "farmlands.farmlandname","crops.cropsname","machines.machine_num"}).
			Joins("JOIN possessions ON possessions.agrist_id = agrists.id").
			Joins("JOIN farmlands ON possessions.farmland_id = farmlands.id").
			Joins("JOIN plantings").
			Joins("JOIN crops ON crops.id = plantings.crops_id").
			Joins("JOIN machines ").
			Where("agrists.username = ?", username).
			Where("farmlands.id = plantings.farmland_id").
			Where("farmlands.id = machines.farmland_id").
			Scan(&farmaland)
			ctx.JSON(http.StatusOK, farmaland)
		}
	}

}