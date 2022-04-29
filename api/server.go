package api

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"gorm.io/gorm"
)

var db *gorm.DB = setupDB()

func mainRoute(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

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
	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	if logging {
		app.Use(logger.New())
	}

	app.Get("/leaderboard", leaderboardRoute)
	app.Get("/", mainRoute)
	app.Post("/", postFlagRoute)
	return app
}
