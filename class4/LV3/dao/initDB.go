package dao

import (
	"database/sql"
	"log"
)

var MyDb *sql.DB

func InitDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	//同步数据
	MyDb = db

	log.Println("DB connect success")
	return
}
