package main

import (
	"GoServer/ent"
	"context"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

func main() {
	client, err := ent.Open("mysql", "root:1234@tcp(127.0.0.1:3316)/go?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("database schema created")

	app := fiber.New()

	app.Post("/make", func(c fiber.Ctx) error {
		var data map[string]interface{}
		json.Unmarshal(c.Body(), &data)
		title, _ := data["title"].(string)
		description, _ := data["description"].(string)
		newTodo, err := client.Todos.
			Create().
			SetTitle(title).
			SetDescription(description).
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating todo ", err)
		}
		log.Printf("User created: %+v\n", newTodo)
		return nil
	})

	app.Get("/:value", func(c fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("value"))
		todo, err := client.Todos.Get(ctx, id)
		if err != nil {
			log.Printf("User not found")
		}
		return c.JSON(fiber.Map{
			"data": todo,
		})
	})

	app.Listen(":3000")
}
