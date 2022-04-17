package api

import (
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

	if err := u.Validate(flag); err != nil {
		return err
	}

	result := db.FirstOrCreate(u)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(u)
}

func SetupRoutes(logging bool) *fiber.App {
	app := fiber.New()

	if logging {
		app.Use(logger.New())
	}

	app.Get("/leaderboard", leaderboardRoute)
	app.Post("/", postFlagRoute)
	return app
}
