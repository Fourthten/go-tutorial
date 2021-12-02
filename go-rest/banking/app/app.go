package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-rest/banking/domain"
	"go-rest/banking/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined....")
	}
}

func Start() {

	sanityCheck()

	// mux := http.NewServeMux() #If not using gorila mux like a router
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	// 	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// 	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// 	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	// => go run main.go
	// log.Fatal(http.ListenAndServe("localhost:8000", router))

	// => SERVER_ADDRESS=localhost SERVER_PORT=8000 go run main.go
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request received")
// }
//
// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
