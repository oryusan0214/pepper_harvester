package get

import (
	"main/entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWeekPhoto(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var Data []entities.DayPhoto
		var OutputData []entities.DayPhoto
		err := ctx.ShouldBindJSON(&Data)
		farmlandname := ctx.Param("name")
		day := ctx.Param("day")
		timelauout := "2006-01-02"
		var incrementDay []string

		println(day)
		indexDay , _:= time.Parse(timelauout,day)
		println(indexDay.String())

		if err != nil {
			for count := 1; count <= 7; count++ {
				incrementDay = []string{strconv.Itoa(indexDay.Year())+"-"+strconv.Itoa(int(indexDay.Month()))+"-"+strconv.Itoa(indexDay.Day())}
				println("count=",count,incrementDay)

				db.Debug().
				Table("farmlands").
				Select([]string{"photos.url"}).
				Joins("JOIN machines ON machines.farmland_id = farmlands.id").
				Joins("JOIN harvests ON harvests.machine_id = machines.id").
				Joins("JOIN photos ON photos.id = harvests.photo_id").
				Where("farmlands.farmlandname = ?", farmlandname).
				Where("harvests.day = ?",incrementDay).
				Scan(&Data)
				
				if Data != nil  {
					if OutputData == nil {
						OutputData = Data
						}	else {
							OutputData = append(OutputData,Data...)
						}
				}
				println("add before=",indexDay.String())
				indexDay=indexDay.AddDate(0,0,1)
				println("add after=",indexDay.String())
			}
			ctx.JSON(http.StatusOK, Data)
		}
	}
}