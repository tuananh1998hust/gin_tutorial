package main

import (
	"github.com/tuananh1998hust/gin_tutorial/config"
	"github.com/tuananh1998hust/gin_tutorial/routes"
)

func main() {
	r := routes.SetUpRoutes()

	// Connect DB
	config.Connect()

	r.Run()
}
