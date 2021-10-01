package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":"+os.Getenv("PORT")))
	log.Println("sukses konek")
}