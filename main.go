package main

import (
	"flag"
	"log"
	"todo/config"
	"todo/database"
	"todo/module"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Demo CSV API
// @version 1.0
// @description Language: Golang. Core: Fiber

// @contact.name CubeSystem Viet Nam
// @contact.url https://vn-cubesystem.com/
// @contact.email info@vn-cubesystem.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-csv-key
// @securityDefinitions.apikey ApiTokenAuth
// @in header
// @name x-csv-token
func main() {
	app := fiber.New()

	// Initialize default config
	app.Use(cors.New())

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
