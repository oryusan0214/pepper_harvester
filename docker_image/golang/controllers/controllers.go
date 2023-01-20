package controllers

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnect()  *gorm.DB{
	err:=godotenv.Load("../.env")
	if err !=nil {
		fmt.Printf("環境ファイルの読み込みに失敗しました。:%v",err)
	}
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	//DBに接続
	db,err := gorm.Open(mysql.Open(CONNECT),&gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DBに接続できました。")
	return db
}