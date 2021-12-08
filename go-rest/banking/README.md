## Source
```bash
https://github.com/ashishjuyal/banking

SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD= DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go
```

## API
```python
# Content-Type (application/xml, application/json)
# GET http://localhost:8000/customers?status=active
# GET http://localhost:8000/customers/{custId}
# POST http://localhost:8000/customers/{custId}/account
    {
        "account_type": "saving", /* saving or checking */
        "amount": 5000.05
    }
# POST http://localhost:8000/customers/{custId}/account/{accId}
    {
        "transaction_type": "deposit", /* deposit or withdrawal */
        "amount": 10000
    }
```

Run `./start.sh` to download the dependencies and run the the application

* option 1: `resources/database.sql` if you don't want to use docker, please import it.
* option 2: configuration the docker file as you wish. to start the docker container, run the `docker-compose up` inside the `resources/docker` folder.