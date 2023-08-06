package database

import (
	"fmt"
	"strconv"
	"time"
	"todo/config"

	"github.com/gofiber/storage/postgres"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var Store *postgres.Storage

func Connect() bool {
	var status bool = true
	var err error
	dbHost := config.Get("DB_HOST")
	dbPort := config.Get("DB_PORT")
	dbUser := config.Get("DB_USER")
	dbPassword := config.Get("DB_PASSWORD")
	dbName := config.Get("DB_NAME")
	dbSSH := config.Get("DB_SSH")
	dbTimezone := config.Get("APP_TIME_ZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSSH, dbTimezone)

	DB, err = gorm.Open(postgresDriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		status = false
		fmt.Println("failed to connect database")
	}

	// connection pool
	dbSetMaxIdleConns, _ := strconv.Atoi(config.Get("DB_SET_MAX_IDLE_CONNS"))
	dbSetMaxOpenConns, _ := strconv.Atoi(config.Get("DB_SET_MAX_OPEN_CONNS"))
	dbSetConnMaxLifetime, _ := time.ParseDuration(config.Get("DB_SET_CONN_MAX_LIFETIME"))

	poolConn, err := DB.DB()

	if err != nil {
		status = false
		fmt.Println("failed to connect database")
	}
	poolConn.SetMaxIdleConns(dbSetMaxIdleConns)
	poolConn.SetMaxOpenConns(dbSetMaxOpenConns)
	poolConn.SetConnMaxLifetime(dbSetConnMaxLifetime)

	ConfigSession()

	return status
}

func ConfigSession() {

	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	user := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	name := config.Get("DB_NAME")
	post, _ := strconv.Atoi(port)

	Store = postgres.New(postgres.Config{
		Host:     host,
		Port:     post,
		Username: user,
		Password: password,
		Database: name,
		Table:    "session",
		Reset:    false,
	})

}
