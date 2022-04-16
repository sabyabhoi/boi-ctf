package main

func main() {
	app := setupRoutes(true)
	app.Listen(":8080")
}
