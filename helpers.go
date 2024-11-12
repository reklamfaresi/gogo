package main

import (
	"log"
)

func RecordActivity(username, action string) {
	query := `INSERT INTO activities (username, action) VALUES (?, ?)`
	_, err := DB.Exec(query, username, action)
	if err != nil {
		log.Println("Aktivite kaydedilemedi:", err)
	}
}
