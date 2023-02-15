package gateways

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
// =====================================
// 	struct
// =====================================
type Database_mysql struct {}


// ######################################
// content:
// インターフェースからメソッドを動かす処理
// ######################################
// ***************************************
// cleate   : ryunosuke yamada
// argument : GatewaysInterface
// return   : Mysqlのオープン処理 *gorm.DB
// ***************************************
func DbConnect( g GatewaysInterface ) *gorm.DB{
	result := g.Connect()
	return result
}


// ######################################
// content :
// 環境ファイルの読み込み。
// ###################################### 
// ***************************************
// cleate   : ryunosuke yamada
// argument : none
// return   : sql.Open()に必要な情報 string
// ***************************************
func GetEnv() string {
	//環境ファイルの読み込み
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("環境ファイルの読み込みに失敗しました。:%v", err)
	}
	//環境変数の読み込み
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	return CONNECT
}

// ######################################
// content :
// sqlへの接続処理。
// ###################################### 
// ***************************************
// cleate   : ryunosuke yamada
// argument : none
// return   : *gorm.DB
// ***************************************
func(db *Database_mysql) Connect() *gorm.DB{
	//環境ファイルから設定情報を取り出す。
	CONNECT := GetEnv()
	//DBに接続
	sqlDB,err := sql.Open("mysql",CONNECT)
	gormDB,err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}),&gorm.Config{})
	//エラー処理
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DBに接続できました。")
	return gormDB
}


