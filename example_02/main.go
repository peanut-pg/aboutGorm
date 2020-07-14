package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Profile      Profile `gorm:"ForeignKey:ProfileRefer"` // 使用ProfileRefer作为外键
	ProfileRefer int
	Name         string
}

type Profile struct {
	gorm.Model
	Name string
}

func InitData(db *gorm.DB) {
	profile := Profile{
		Name: "first profile",
	}
	// 初始化一些数据
	db.Create(&profile)
	profile2 := Profile{}
	db.First(&profile2, "name=?", profile.Name)
	user := &User{
		Profile: profile2,
		Name:    "贩子",
	}
	db.Create(user)
}

func getProfile(db *gorm.DB) {
	profile2 := Profile{}
	user := User{}
	// db.First 可以用于查询数据
	db.First(&user, "name=?", "贩子")
	// 因为我们上面用`gorm:"ForeignKey:ProfileRefer"`手动指定了外键，
	// 所以这里的Related的第二个参数也需要手动指定
	db.Model(&user).Related(&profile2, "profile_refer")
	fmt.Println(profile2)

}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(10.211.55.3:3306)/aboutGorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.AutoMigrate(&Profile{}, &User{})
	// 开启日志模式之后可以看到详细的sql语句执行
	db.LogMode(true)
	InitData(db)
	getProfile(db)
}
