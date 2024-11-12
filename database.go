package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gogo")
	if err != nil {
		log.Fatal("Veritabanı bağlantısı sağlanamadı:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Veritabanına ulaşılamadı:", err)
	}

	log.Println("Veritabanı bağlantısı başarılı!")
}
