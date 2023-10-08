package main

import (
	"os"
	"otus-social-network/internal/app"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logrus.SetLevel(ll)
}

func main() {
	a, err := app.NewApp()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = a.Run()
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
