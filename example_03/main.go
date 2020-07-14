package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	Name      string
	User      User `gorm:"foreignkey:UserRefer"`
	UserRefer uint
}

func InitData(db *gorm.DB) {
	user := User{
		Name: "贩子",
	}
	db.Create(&user)

	profile := &Profile{
		Name: "first profile",
		User: user,
	}
	db.Create(&profile)
}

func getProfile(db *gorm.DB) {
	profile2 := Profile{}
	user := User{}
	// db.First 可以用于查询数据
	db.First(&user, "name=?", "贩子")
	db.Model(&user).Related(&profile2, "user_refer")
	fmt.Println(profile2)
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(10.211.55.3:3306)/aboutGorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&Profile{}, &User{})
	//InitData(db)
	getProfile(db)
}
