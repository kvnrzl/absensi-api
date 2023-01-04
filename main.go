package main

import "log"

func main() {
	r := InitServer()
	log.Fatal(r.Run("127.0.0.1:8080"))
}
