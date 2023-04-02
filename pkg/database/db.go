package database

import (
	"fmt"

  moviePb "github.com/ahnafms/crud-grpc-go/pkg/protobuf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectDB() {
	USER := "root"
	PASS := "root"
	HOST := "localhost"
	PORT := 3306
	DBNAME := "pi_grpc"
	URL := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	database, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = database
  DB.AutoMigrate(&moviePb.Movie{})
}
