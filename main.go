package main

import "github.com/sabyabhoi/boi-ctf/api"

func main() {
	app := api.SetupRoutes(true)
	app.Listen(":8080")
}
