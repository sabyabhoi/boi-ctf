package main

func main() {
	app := setupRoutes()
	app.Listen(":8080")
}
