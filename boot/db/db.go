package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"
)

var DB *sql.DB

func InitRedis() {

}

func InitMysql() {
	dataSourceName := fmt.Sprintf("root:@tcp(localhost:3306)/?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("连接数据库失败")
		_ = db.Close()
		panic(err)
	}
	DB = db
}

func InitDb() {
	InitMysql()
	InitRedis()

}
