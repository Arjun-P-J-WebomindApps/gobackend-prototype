package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Connect(host string,user string,password string,dbName string,dbPort string,sslMode string) {
	dsn:=fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbName,
		dbPort,
		sslMode,
	)

	database,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	

	if err!=nil{
		panic("Error connecting to database"+err.Error())
	}

	DB=database
}