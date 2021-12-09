## Getting Started
```
go run main.go
```

## Route API
`http://localhost:3000`,

Create ToDo:
```
{
    "name": "names",
    "todo": "This is To Do"
}
```
Get By Name: `http://localhost:3000?name=names`
Delete: `http://localhost:3000/names`

| Task | URL | Method | Response code | Response |
|:----:|:---:|:------:|:-------------:|:--------:|
| Check API health | /ping | GET | 200 | pong |
| Create an entry | / | POST | 201 | json of the created object | 
| Read all entries | /all | GET | 200 | json of all entries |
| Read entries by id | /?id=value | GET | 200 | json of entries using keys and values | 
| Delete entries | /:id | DELETE | 200 | status |

## Source
```bash
https://github.com/go-sql-driver/mysql
go get -u github.com/go-sql-driver/mysql

https://github.com/golang/go/wiki/Modules
```

## Database
```
CREATE DATABASE mysqldb;
CREATE TABLE TODO(name varchar(50), todo varchar(1000));
```

## Heroku
This is a simple deploy to cloud.
first, add port in `main.go`.
```
git init
git add .
git commit -m "first commit"

brew tap heroku/brew && brew install heroku
heroku login
heroku create nameof-apps
heroku git:remote -a nameof-apps
git push heroku master

https://dashboard.heroku.com/apps
https://nameof-apps.herokuapp.com/ping
```