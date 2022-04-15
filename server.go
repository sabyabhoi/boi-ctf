package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

var db *gorm.DB = setupDB()

func leaderboardRoute(c *fiber.Ctx) error {

	var users []User

	db.Order("created_at").Find(&users)

	return c.JSON(users)
}

func postFlagRoute(c *fiber.Ctx) error {
	c.Accepts("application/json")

	u := new(User)

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	u.CreatedAt = time.Now()

	flag := os.Getenv("FLAG")
	log.Println(flag)

	if err := u.Validate(flag); err != nil {
		return err
	}

	return c.JSON(u)
}

func setupServer(port string) {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/leaderboard", leaderboardRoute)
	app.Post("/", postFlagRoute)

	app.Listen(port)
}
