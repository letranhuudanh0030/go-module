package main

import (
	"flag"
	"log"
	"todo/config"
	"todo/database"
	"todo/module"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if !database.Connect() {
		panic("Fail connection database")
	}

	module.AutoMigrate()
	module.InitRoute(app)

	port := config.Get("ENV_PORT")

	addr := flag.String("addr", ":"+port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
