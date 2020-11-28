package main

import (
	"github.com/prasetiyo28/card-api/db"
	"github.com/prasetiyo28/card-api/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":3000"))
}
