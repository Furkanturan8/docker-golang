package main

import (
	"context"
	"docker-golang/book"
	"docker-golang/database"
	"log"
	"time"
)

func main() {
	db, err := database.NewMysqlDB(database.MysqlConfig{
		Username: "your-username",
		Password: "your-password",
		DbName:   "your-dbName",
		Port:     3307,
		Host:     "localhost",
	})
	if err != nil {
		log.Fatal("impossible to create mysql db : %w", err)
	}

	b, err := db.Create(context.Background(), book.Book{
		Name:       "İrade Eğitimi",
		AuthorName: "Jules Payot",
		CreateTime: time.Now(),
	})
	if err != nil {
		log.Fatalf("impossible to insert: %s", err)
	}
	log.Println(b)
}
