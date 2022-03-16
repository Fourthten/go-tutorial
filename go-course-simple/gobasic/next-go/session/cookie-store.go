package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

const SESSION_ID = "id"

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true
	return store
}

func main() {
	store := newCookieStore()
	e := echo.New()
	e.Use(echo.WrapMiddleware(context.ClearHandler))
	e.GET("/set", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Values["message1"] = "Hello"
		session.Values["message2"] = "World"
		session.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})
	e.GET("/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}
		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})
	e.GET("/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})
	e.Logger.Fatal(e.Start(":9000"))
}

// ArangoDB (https://github.com/starJammer/gorilla-sessions-arangodb)
// Bolt (https://github.com/yosssi/boltstore)
// Couchbase (https://github.com/srinathgs/couchbasestore)
// DynamoDB on AWS (https://github.com/denizeren/dynamostore)
// DyanmoDB on AWS (Official AWS Library) (https://github.com/savaki/dynastore)
// Memcache (https://github.com/bradleypeabody/gorilla-sessions-memcache)
// Memcache/Datastore/Context in AppEngine (https://github.com/dsoprea/go-appengine-sessioncascade)
// MongoDB (https://github.com/kidstuff/mongostore)
// MySQL (https://github.com/srinathgs/mysqlstore)
// MySQL Cluster (https://github.com/EnumApps/clustersqlstore)
// PostgreSQL (https://github.com/antonlindstrom/pgstore)
// Redis (https://github.com/boj/redistore)
// RethinkDB (https://github.com/boj/rethinkstore)
// Riak (https://github.com/boj/riakstore)
// SQLite (https://github.com/michaeljs1990/sqlitestore)
// GORM (MySQL, PostgreSQL, SQLite) (https://github.com/wader/gormstore)
// ql (https://github.com/gernest/qlstore)
// In-memory implementation for use in unit tests (https://github.com/quasoft/memstore)
// XORM (MySQL, PostgreSQL, SQLite, Microsoft SQL Server, TiDB) (https://github.com/lafriks/xormstore)
