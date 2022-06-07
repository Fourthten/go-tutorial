package databases

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	// Min connection can create
	db.SetMaxIdleConns(10)
	// Max connection can create
	db.SetMaxOpenConns(100)
	// Connection time not use can delete
	db.SetConnMaxIdleTime(5 * time.Minute)
	// Connection time can use
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
