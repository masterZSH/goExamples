package main

import "fmt"
import  "github.com/jinzhu/gorm"

import (
    _ "github.com/go-sql-driver/mysql"
)


type Site struct {
	ID        uint   `gorm:"primary_key;column:Id"`
	SiteName  string `gorm:"column:site_id"`
	SiteURL   string `gorm:"column:site_url"`
	SiteFtp   string `gorm:"column:site_ftp"`
	SiteAdmin string `gorm:"column:site_admin"`
}

func main() {
	var site Site
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/web?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)

	// 查询第一个
	db.First(&site)
	fmt.Print(err)
	fmt.Printf("%+v",site)
    err=db.Delete(&site).Error
    if err !=nil{
        fmt.Println(err)
    }


	defer db.Close()
}