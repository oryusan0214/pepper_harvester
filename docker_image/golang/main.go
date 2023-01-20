package main

import (
	"log"
	"main/controllers"
	"main/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var agrigator entities.Agrist
	var possession entities.Possession
	var farmland entities.Farmland
	var db *gorm.DB

	db = controllers.DbConnect()
	//関数登録
	r := gin.Default()
	r.GET("/Agrist",func(ctx *gin.Context) {
		err:=ctx.ShouldBindJSON(&agrigator)
		//なおエラーの時はAbortWithErrorメソッドが呼ばれるのでエラー処理が自動で行われる

		if err == nil{
			result :=db.Where( "name <> ?", agrigator.Name).Find(&agrigator)
			result =db.Where("id <> ?",agrigator.Possession_id).Find(&possession)
			result =db.Where("id <> ?",possession.Farmland_id).Find(&farmland)
			
			var msg struct{
				name string 
				farmname string 
			}
			msg.name=agrigator.Name
			msg.farmname=farmland.Name
			println(msg.name)
			println(msg.farmname)

			if  result.Error !=nil {
				println(result)
			}
			result = db.Table("agrist").
			Select("agrist.name,farmland.name").Joins("JOIN possession ON agrist.id=possession.id").Joins("JOIN farmland ON possession.farmland_id=farmland.id").Select(&msg)
			if result != nil{
				println(result)	
			}
			

			
			ctx.JSON(http.StatusOK,msg)
		}
	})
	
	err := r.Run(":8080")
	if err !=nil{
		log.Fatalf("impossible to start server: %s",err)
	}
	
	//DBをクローズ
	dbClose, err := db.DB()
	defer dbClose.Close()	
	
}

