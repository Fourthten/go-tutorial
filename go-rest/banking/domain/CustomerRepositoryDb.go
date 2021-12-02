package domain

import (
	"database/sql"
	"go-rest/banking/errs"
	"go-rest/banking/logger"

	"github.com/jmoiron/sqlx"

	// "log"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	// client *sql.DB
	client *sqlx.DB
}

// func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
// 	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

// 	rows, err := d.client.Query(findAllSql)
// 	if err != nil {
// 		log.Println("Error while querying customer table " + err.Error())
// 		return nil, err
// 	}

// 	customers := make([]Customer, 0)
// 	for rows.Next() {
// 		var c Customer
// 		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
// 		if err != nil {
// 			log.Println("Error while scanning customers " + err.Error())
// 			return nil, err
// 		}
// 		customers = append(customers, c)
// 	}
// 	return customers, nil
// }

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		// rows, err = d.client.Query(findAllSql, status)
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		// log.Println("Error while querying customer table " + err.Error())
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		// log.Println("Error while scanning customers " + err.Error())
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// row := d.client.QueryRow(customerSql, id)
	var c Customer
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			// log.Panicln("Error while scanning customer " + err.Error())
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// func NewCustomerRepositoryDb() CustomerRepositoryDb {
// 	// client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
// 	// client, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")

// 	dbUser := os.Getenv("DB_USER")
// 	dbPasswd := os.Getenv("DB_PASSWD")
// 	dbAddr := os.Getenv("DB_ADDR")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbName := os.Getenv("DB_NAME")

// 	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
// 	client, err := sqlx.Open("mysql", dataSource)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// See "Important settings" section.
// 	client.SetConnMaxLifetime(time.Minute * 3)
// 	client.SetMaxOpenConns(10)
// 	client.SetMaxIdleConns(10)
// 	return CustomerRepositoryDb{client}
// }

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
