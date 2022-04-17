package main

import "github.com/sabyabhoi/boi-ctf/api"

func main() {
	go StartSocat("8082","cat -")
	app := api.SetupRoutes(true)
	app.Listen(":8081")
}
