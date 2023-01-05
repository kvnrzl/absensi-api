package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	r := InitServer()
	logrus.Info("Server started at port 8080")
	logrus.Fatal(r.Run(":8080"))
}
