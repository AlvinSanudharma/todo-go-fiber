package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func initDb() (*sql.DB, error) {
	dns := "user=postgres.vqxvgfptkbsxfkzbiiti password=wddRRgS9Li97oZcG host=aws-1-ap-northeast-2.pooler.supabase.com port=6543 dbname=postgres"

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
