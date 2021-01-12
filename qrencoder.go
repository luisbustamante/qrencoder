package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	port := getenv("PORT", "3000")
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(":" + port))
}
func setupRoutes(app *fiber.App) {
	app.Get("/generate", Generate)
}

//Generate returns qr code as byte[] image
//example: /generate?text=example.com&size=256
func Generate(c *fiber.Ctx) error {
	var png []byte
	msg := fmt.Sprintf("%s", c.Query("text"))
	if len(msg) == 0 {
		return c.Status(400).SendString("Must specify text to encode")
	}
	size := fmt.Sprintf("%s", c.Query("size"))
	qrsize, sizeErr := strconv.Atoi(size)
	if sizeErr != nil {
		qrsize = 256
	}
	png, err := qrcode.Encode(msg, qrcode.Medium, qrsize)
	if err != nil {
		return c.Status(500).SendString("Error")
	}
	c.Set("Content-Type", "image/png")
	return c.Send(png)
}
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
