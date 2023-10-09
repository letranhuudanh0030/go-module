package main

import (
	"flag"
	"fmt"
	"log"
	"todo/config"
	"todo/database"
	"todo/module"
	"todo/module/moduleA"
	"todo/module/moduleB"

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

	// Test
	database.DB.AutoMigrate(&moduleA.User{}, &moduleB.Order{}, &User{}, &Role{})

	database.DB.Migrator().CreateConstraint(&moduleB.Order{}, "User")
	database.DB.Migrator().CreateConstraint(&moduleB.Order{}, "fk_orders_users")

	userService := moduleA.NewUserServiceImpl()
	orderService := moduleB.NewOrderServiceImpl(userService)

	orderID := 123
	orderWithUser := orderService.GetOrderByID(orderID)

	fmt.Printf("General: %+v\n", orderWithUser)
	fmt.Printf("Order: %+v\n", orderWithUser.Order)
	fmt.Printf("User: %+v\n", orderWithUser.User)

	// Test many to many
	user := User{}
	role := Role{}
	database.DB.First(&user, 1)
	database.DB.First(&role, 1)
	database.DB.Model(&user).Association("Roles").Append(&role)

	// user := User{}
	// db.Preload("Roles").First(&user, 1)

	// Run App
	port := config.Get("ENV_PORT")

	addr := flag.String("addr", ":"+port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}

type User struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Users []User `gorm:"many2many:user_roles;"`
}

type UserRole struct {
	UserID uint
	RoleID uint
	Desc   string
}
