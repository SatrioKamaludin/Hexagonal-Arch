package main

import (
	"CRUD-Go-Hexa-MongoDB/internal/app"
	"log"
)

func main() {
	app := app.Setup()
	log.Fatal(app.Listen(":3000"))
}
