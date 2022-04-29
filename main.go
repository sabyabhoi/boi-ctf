package main

import "github.com/sabyabhoi/boi-ctf/api"

func main() {
	go StartSocat("8081","./run.sh")
	app := api.SetupRoutes(true)
	app.Listen(":8080")
}
