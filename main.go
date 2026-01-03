package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDb() (*sql.DB, error) {
	dns := os.Getenv("DATABASE_URL")
	if dns == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", dns)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()

		return nil, err
	}

	return  db, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using OS env")
	}
	
	db, err := initDb()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Go Fiber!")
	})

	app.Listen(":8182") // TODO: CHANGE TO 8081 LATER !
}
