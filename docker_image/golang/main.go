package main

import (
	"log"
	"main/gateways"
	"main/get"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	//変数宣言
	var db *gorm.DB
	var gateway gateways.GatewaysInterface
	var database gateways.Database_mysql
	gateway=&database
  //mysqlに接続
	db = gateways.DbConnect(gateway)

	//DBをクローズ
	 dbClose, err := db.DB()
	 defer dbClose.Close()

	// 関数登録
	// ginを作成
	router := gin.Default()
	
	router.GET("/Agrist/:username",get.GetUser(db))
	router.GET("/farmland/:username",get.GetFarmland(db))

	err = router.Run(":8080")
	if err !=nil{
		log.Fatalf("impossible to start server: %s",err)
	}
}

