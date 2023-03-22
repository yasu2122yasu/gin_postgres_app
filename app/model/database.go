package model

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "host=db port=5432 user=postgres password=postgres dbname=app sslmode=disable"
	dialector := postgres.Open(dsn)
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}
