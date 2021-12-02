## Source
```bash
https://github.com/ashishjuyal/banking
https://github.com/ashishjuyal/banking-auth
https://github.com/ashishjuyal/banking-lib

https://github.com/gorilla/mux
go get -u github.com/gorilla/mux

https://github.com/go-sql-driver/mysql
go get -u github.com/go-sql-driver/mysql

https://github.com/uber-go/zap
go get -u go.uber.org/zap

https://github.com/jmoiron/sqlx
go get github.com/jmoiron/sqlx

SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD= DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go
```

## API
```python
# Content-Type (application/xml, application/json)
# GET http://localhost:8000/customers?status=active
# GET http://localhost:8000/customers/2000
# POST http://localhost:8000/customers/2000/account
    {
        "account_type": "saving",
        "amount": "5000.05"
    }
```

[REST based microservices API development in Golang](https://www.udemy.com/share/103O5K3@f7Pf9EPs_ILg-vPmCF-CuzvjI8WaFfH85UbN5BCNxZO7cIVE9Q9Cz-l1flcV1u0Q/)