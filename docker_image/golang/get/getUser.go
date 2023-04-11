package get

import (
	"main/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func  GetUser(db *gorm.DB )func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		var agrigator []entities.Agrist
	
		err:=ctx.ShouldBindJSON(&agrigator)
		username := ctx.Param("username")
		
		if err != nil{
			  db.Debug().
				Table("agrists").
				Select([]string{"agrists.id","agrists.username","farmlands.id","farmlands.farmlandname"}).
				Joins("JOIN possessions ON possessions.agrist_id = agrists.id").
				Joins("JOIN farmlands ON possessions.farmland_id = farmlands.id").
				Where("agrists.username = ?",username).
				Scan(&agrigator)
				
				ctx.JSON(http.StatusOK,agrigator)
		}
	}
	
}