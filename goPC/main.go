package main

import (
	"fmt"
	"goPC/dao"
	"goPC/model"
	"goPC/router"
)

// Admin 管理员
func main() {
	//dsn := "root:111111@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql success")
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&model.Admin{})
	dao.DB.AutoMigrate(&model.Advise{})
	dao.DB.AutoMigrate(&model.Room{})
	dao.DB.AutoMigrate(&model.User{})
	dao.DB.AutoMigrate(&model.Reserve{})
	r := router.SetupRouter()
	r.Run(":8090")
}
