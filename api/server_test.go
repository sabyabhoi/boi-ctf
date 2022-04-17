package api

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = SetupRoutes(false)

func TestLeaderboardRoute(t *testing.T) {
	req := httptest.NewRequest("GET", "/leaderboard", nil)

	app.Test(req)
}

func TestPostFlagRoute(t *testing.T) {
	u := []byte(fmt.Sprintf(`{"ID": "2020B3A72147H", "flag": "%s"}`, os.Getenv("FLAG")))

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(u))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, fiber.StatusOK)
}
