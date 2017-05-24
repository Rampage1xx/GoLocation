package Models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)



type DataObject struct {
	gorm.Model
	Longitude string
	Latitude string
	DownloadedAt string
	Country string
	Age string
	Sex string

}
type User struct {
	name string
	password string
}


func Main() *gorm.DB  {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=meeseeks dbname=pinit password=MEESEEKS sslmode=disable")
	//defer db.Close()
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(false)
	db.CreateTable(&DataObject{})

	return db


}


var DbSession = Main()