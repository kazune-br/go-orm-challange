package gormorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func RunGorm() {
	// initialize db
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		"root",
		"password",
		"localhost",
		"4401",
		"gormdb",
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Printf("fail to initialize db")
		panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		log.Printf("fail to initialize db")
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Executing Gorm Example")

	rand.Seed(time.Now().UnixNano())
	i := uint64(rand.Intn(1000))
	u := User{UID: i}
	db.Create(&u)

	users := []User{}
	db.Find(&users)
	fmt.Printf("%v\n", users)

	db.Where("UID = ?", i).Delete(&User{})
}

type User struct {
	// gorm.Model
	// https://gorm.io/ja_JP/docs/conventions.html
	ID  uint64 `gorm:"primary_key"`
	UID uint64 `json:"uid"`
}