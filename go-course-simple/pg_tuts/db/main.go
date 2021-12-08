package db

import (
	"log"
	"os"
	"time"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:         "postgres",
		Password:     "",
		Addr:         "localhost:5432",
		Database:     "tuts",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		MaxConnAge:   1 * time.Minute,
		PoolSize:     20,
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connection to database successful.\n")
	CreateProdItemsTable(db)
	return db

	// closeErr := db.Close()
	// if closeErr != nil {
	// 	log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
	// 	os.Exit(100)
	// }
	// log.Printf("Connection closed successfully")
	// return
}
