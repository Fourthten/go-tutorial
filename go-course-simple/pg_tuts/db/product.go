package db

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type ProductItem struct {
	RefPointer int      `sql:"-"`
	tableName  struct{} `sql:"product_items_collection"`
	ID         int      `sql:"id,pk"`
	Name       string   `sql:"name,unique"`
	Desc       string   `sql:"desc"`
	Image      string   `sql:"image"`
	Price      float64  `sql:"price,type:real"`
	Features   struct {
		Name string
		Desc string
		Imp  int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

func (pi *ProductItem) GetById(db *pg.DB) error {
	// getErr := db.Select(pi)
	// getErr := db.Model(pi).Column("name", "desc").Where("id = ?0", pi.ID).Select()
	var items []ProductItem
	getErr := db.Model(&items).Column("name", "desc").
		Where("id = ?0", pi.ID).
		WhereOr("id = ?0", 2).
		Offset(0).
		Limit(2).
		Order("id asc").
		Select()

	if getErr != nil {
		log.Printf("Error while getting value by id, Reason: %v\n", getErr)
		return getErr
	}
	// log.Printf("Get by id succesfull for %v\n", *pi)
	log.Printf("Get by id succesfull for %v\n", items)
	return nil
}

func (pi *ProductItem) UpdatePrice(db *pg.DB) error {
	tx, txErr := db.Begin()
	if txErr != nil {
		log.Printf("Error while opening tx, Reason: %v\n", txErr)
		return txErr
	}

	_, updateErr := tx.Model(pi).Set("price = ?price").Where("id = ?id").Update()
	// _, updateErr := db.Model(pi).Set("price = ?price").Where("id = ?id").Update()
	if updateErr != nil {
		log.Printf("Error while updating price. Reason: %v\n", updateErr)
		tx.Rollback()
		return updateErr
	}
	log.Printf("%v\n", pi)
	_, updateErr2 := tx.Model(pi).Set("is_active = ?0", false).Where("id = ?id").Update()
	if updateErr2 != nil {
		log.Printf("Error while updating is_active. Reason: %v\n", updateErr2)
		tx.Rollback()
		return updateErr
	}

	tx.Commit()
	log.Printf("Price updated successfully for ID %d\n", pi.ID)
	return nil
}

func (pi *ProductItem) DeleteItem(db *pg.DB) error {
	_, deleteErr := db.Model(pi).Where("name = ?name").Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting items, Reason: %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Delete successful for %s, Item\n", pi.Name)
	return nil
}

func (pi *ProductItem) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("ProductItem %s inserted successfully.\n", pi.Name)
	return nil
}

func (pi *ProductItem) SaveAndReturn(db *pg.DB) (*ProductItem, error) {
	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item, Reason: %v\n", insertErr)
		return nil, insertErr
	}
	log.Printf("ProductItem Inserted successfully.\n")
	log.Printf("Received new value result is: %v\n", InsertResult.RowsAffected)
	return pi, nil
}

func (pi *ProductItem) SaveMultiple(db *pg.DB, items []*ProductItem) error {
	_, insertErr := db.Model(items[0], items[1]).Insert()
	if insertErr != nil {
		log.Printf("Error while inserting bulk items, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Bulk Insert successful\n")
	return nil
}

func CreateProdItemsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table productItems, Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table ProductItems created successfully.\n")
	return nil
}
