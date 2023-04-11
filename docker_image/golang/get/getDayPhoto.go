package get

import (
	"main/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDayPhoto(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var Data []entities.DayPhoto

		err := ctx.ShouldBindJSON(&Data)
		farmlandname := ctx.Param("name")
		day := ctx.Param("day")

		if err != nil {
			db.Debug().
				Table("farmlands").
				Select([]string{"photos.url" }).
				Joins("JOIN machines ON machines.farmland_id = farmlands.id").
				Joins("JOIN harvests ON harvests.machine_id = machines.id").
				Joins("JOIN photos ON photos.id = harvests.photo_id").
				Where("farmlands.farmlandname = ?", farmlandname).
				Where("harvests.day = ?", day).
				Scan(&Data)
			ctx.JSON(http.StatusOK, Data)
		}
		
	}

}