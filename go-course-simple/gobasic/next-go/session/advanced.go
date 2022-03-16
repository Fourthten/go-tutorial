package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/kidstuff/mongostore"
	"github.com/labstack/echo"
)

// session
type SessionManager struct {
	store    sessions.Store
	valueKey string
}

func NewSessionManager(store sessions.Store) *SessionManager {
	s := new(SessionManager)
	s.valueKey = "data"
	s.store = store
	return s
}

func (s *SessionManager) Get(c echo.Context, name string) (interface{}, error) {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}
	if val, ok := session.Values[s.valueKey]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}

func (s *SessionManager) Set(c echo.Context, name string, value interface{}) error {
	session, _ := s.store.Get(c.Request(), name)
	session.Values[s.valueKey] = value
	err := session.Save(c.Request(), c.Response())
	return err
}

func (s *SessionManager) Delete(c echo.Context, name string) error {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}

// main
var sessionManager *SessionManager

const SESSION_ID = "id"

type UserModel struct {
	ID   string
	Name string
	Age  int
}

func newMongoStore() *mongostore.MongoStore {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	dbCollection := session.DB("learngo").C("session")
	maxAge := 86400 * 7
	ensureTTL := true
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")
	store := mongostore.NewMongoStore(dbCollection, maxAge, ensureTTL, authKey, encryptionKey)
	return store
}

func main() {
	gob.Register(UserModel{})
	sessionManager = NewSessionManager(newMongoStore())
	e := echo.New()
	e.Use(echo.WrapMiddleware(context.ClearHandler))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	eTest := e.Group("/test")
	eTest.GET("/set", func(c echo.Context) error {
		user := new(UserModel)
		user.ID = "001"
		user.Name = "Setia"
		user.Age = 20
		err := sessionManager.Set(c, SESSION_ID, user)
		if err != nil {
			return err
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/test/get")
	})

	eTest.GET("/get", func(c echo.Context) error {
		result, err := sessionManager.Get(c, SESSION_ID)
		if err != nil {
			return err
		}
		if result == nil {
			return c.String(http.StatusOK, "empty result")
		} else {
			user := result.(UserModel)
			return c.JSON(http.StatusOK, user)
		}
	})

	eTest.GET("/delete", func(c echo.Context) error {
		err := sessionManager.Delete(c, SESSION_ID)
		if err != nil {
			return err
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/test/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
