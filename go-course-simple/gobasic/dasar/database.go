package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type student5 struct {
	id    string
	name  string
	age   int
	grade int
}

var ctx = context.Background()

type student6 struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*sql.DB, error) {
	// user:password@tcp(host:port)/dbname
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_tutorial_golang")
	if err != nil {
		return nil, err
	}
	return db, nil

	// sql.Open("postgres", "user=postgres password=secret dbname=test sslmode=disable")
	// https://github.com/golang/go/wiki/SQLDrivers
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []student5
	for rows.Next() {
		var each = student5{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student5{}
	var id = "E001"
	err = db.
		QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var result1 = student5{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)
	var result2 = student5{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

	stmt2, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
	stmt2.Exec("C001", "Ajung", 21, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stmt3, err := db.Prepare("delete from tb_student where id = ?")
	stmt3.Exec("C001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "A001", "Setia", 20, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 21, "A001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success")

	_, err = db.Exec("delete from tb_student where id = ?", "A001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success")
}

func connect2() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("tutorial_golang"), nil
}

func insert() {
	db, err := connect2()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student6{"Setia", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student6{"Ajung", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success")
}

func find() {
	db, err := connect2()
	if err != nil {
		log.Fatal(err.Error())
	}
	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Setia"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student6, 0)
	for csr.Next(ctx) {
		var row student6
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name: ", result[0].Name)
		fmt.Println("Grade: ", result[0].Grade)
	}
}

func update() {
	db, err := connect2()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Setia"}
	var changes = student6{"Ajung", 2}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Update success")
}

func remove() {
	db, err := connect2()
	if err != nil {
		log.Fatal(err.Error())
	}
	var selector = bson.M{"name": "Ajung"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Remove success")
}

func aggregatedb() {
	db, err := connect2()
	if err != nil {
		log.Fatal(err.Error())
	}

	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": null,
				"Total": { "$sum": 1 }
			} },
			{ "$project": {
				"Total": 1,
				"_id": 0
			} }
		]
	`)), true, &pipeline)
	if err != nil {
		fmt.Println("net")
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("net2")
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)
	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}
	if len(result) > 0 {
		fmt.Println("Total: ", result[0]["Total"])
	}
}

func main() {
	sqlQuery()
	sqlQueryRow()
	sqlPrepare()
	sqlExec()

	insert()
	find()
	update()
	remove()

	aggregatedb()
}
