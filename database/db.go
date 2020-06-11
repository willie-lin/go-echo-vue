package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func Connect() *gorm.DB{
	//db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=true")
	db, err := gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=true")

	if err != nil {
		fmt.Print("データベース接続に失敗しました。")
	}
	db.LogMode(true)
	return db
}
