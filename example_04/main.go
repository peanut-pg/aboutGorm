package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func InitData(db *gorm.DB) {
	credit := CreditCard{
		Number: "12345678900000",
	}
	db.Create(&credit)
	user := User{
		Name:       "贩子",
		CreditCard: credit,
	}
	db.Create(&user)
}

func GetData(db *gorm.DB) {
	user := User{}
	db.First(&user, "name=?", "贩子")
	credit := CreditCard{}
	db.Model(&user).Related(&credit)
	fmt.Println(credit)
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(10.211.55.3:3306)/aboutGorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&CreditCard{}, &User{})
	InitData(db)
	GetData(db)
}
