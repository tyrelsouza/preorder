package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Obviously change this to your settings and get securely.
	//dsn := "root@unix(/var/run/mysqld/mysqld.sock)/db?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN: dsn,
	//}), &gorm.Config{})
	dsn := "root:password@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	_ = db.AutoMigrate(&Author{})
	_ = db.AutoMigrate(&Format{})
	_ = db.AutoMigrate(&Order{})

	DB = db
}
