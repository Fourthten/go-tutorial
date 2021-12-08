package main

import (
	"log"
	"time"

	db "go-course-simple/pg_tuts/db"

	"github.com/go-pg/pg"
)

func main() {
	log.Printf("Hello World...!!!\n")
	pg_db := db.Connect()
	// SaveProduct(pg_db)
	// db.PlaceHolderDemo(pg_db)
	// DeleteItem(pg_db)
	UpdateItemPrice(pg_db)
	// GetById(pg_db)
}

// func SaveProduct(dbRef *pg.DB) {
// 	newPI := &db.ProductItem{
// 		Name:  "Product 2",
// 		Desc:  "Product 2 desc",
// 		Image: "This is image path",
// 		Price: 4.5,
// 		Features: struct {
// 			Name string
// 			Desc string
// 			Imp  int
// 		}{
// 			Name: "F1",
// 			Desc: "F1 Desc",
// 			Imp:  3,
// 		},
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		IsActive:  true,
// 	}
// 	// newPI.Save(dbRef)
// 	newPI.SaveAndReturn(dbRef)
// }

func SaveProduct(dbRef *pg.DB) {
	newPI1 := &db.ProductItem{
		Name:  "Product 3",
		Desc:  "Product 3 desc",
		Image: "This is image path",
		Price: 4.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 Desc",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	newPI2 := &db.ProductItem{
		Name:  "Product 4",
		Desc:  "Product 4 desc",
		Image: "This is image path",
		Price: 4.5,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 Desc",
			Imp:  3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	totalItems := []*db.ProductItem{newPI1, newPI2}
	newPI1.SaveMultiple(dbRef, totalItems)
}

func DeleteItem(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name: "Product 4",
	}
	newPI.DeleteItem(dbRef)
}

func UpdateItemPrice(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID:    1,
		Price: 7.0,
	}
	newPI.UpdatePrice(dbRef)
}

func GetById(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID:   1,
		Name: "Product 1",
	}
	newPI.GetById(dbRef)
}
