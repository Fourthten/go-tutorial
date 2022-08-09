package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "setia").Info("Hello World")

	logger.WithField("username", "ajung").
		WithField("name", "Setia Ajung").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "setia",
		"name":     "Setia Ajung",
	}).Info("Hello World")
}
