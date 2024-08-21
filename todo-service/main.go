package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"todoservice/ent"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := setupDatabase()
	defer db.Close()

	app := fiber.New()

	app.Post("/todos", createTodoHandler(db))
	app.Get("/todos", getTodosHandler(db))

	log.Fatal(app.Listen(":3000"))
}

func getTodosHandler(db *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		todos, err := getTodos(c.Context(), db)
		if err != nil {
			panic(err)
		}
		return c.JSON(todos)
	}
}

func createTodoHandler(db *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload CreateTodoPayload

		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		todo, err := createTodo(c.Context(), db, payload)
		if err != nil {
			panic(err)
		}
		return c.JSON(todo)
	}
}

func setupDatabase() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func createTodo(ctx context.Context, db *ent.Client, todo CreateTodoPayload) (*ent.Todo, error) {
	builder := db.Todo.Create().
		SetDescription(todo.Description)
	if todo.DueAt != nil {
		builder.SetDueAt(*todo.DueAt)
	}

	dbRecord, err := builder.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating todo: %w", err)
	}
	log.Println("todo was created: ", dbRecord)
	return dbRecord, nil
}

func getTodos(ctx context.Context, db *ent.Client) ([]*ent.Todo, error) {
	return db.Todo.Query().All(ctx)
}

type CreateTodoPayload struct {
	Description string     `json:"description"`
	DueAt       *time.Time `json:"due_at"`
}
