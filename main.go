package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	r := InitServer()
	logrus.Info("Server started at port 1312")
	logrus.Fatal(r.Run(":1312"))
}
