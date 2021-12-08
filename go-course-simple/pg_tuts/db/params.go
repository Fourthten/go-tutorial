package db

import (
	"log"

	"github.com/go-pg/pg"
)

type Params struct {
	Param1 string
	Param2 string
}

func PlaceHolderDemo(db *pg.DB) error {
	// var value int
	var value string
	params := Params{
		Param1: "This is param 1",
		Param2: "This is param 2",
	}
	// var query string = "SELECT ?0"
	var query string = "SELECT ?param1"
	// _, selectErr := db.Query(pg.Scan(&value), query, 42, 41)
	_, selectErr := db.Query(pg.Scan(&value), query, params)
	if selectErr != nil {
		log.Print("Error while running the select query, Reason: %v\n", selectErr)
		return selectErr
	}
	log.Printf("Scan successful, Scanned value: %d\n", value)
	return nil
}
