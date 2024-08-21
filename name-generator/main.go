package main

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Name struct {
	ID        string
	FirstName string
	LastName  string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := findRandomName()

		return c.JSON(name)
	})

	log.Fatal(app.Listen(":3000"))
}

func findRandomName() Name {
	return names[rand.Intn(len(names))]
}

var names = []Name{
	{
		ID:        uuid.NewString(),
		FirstName: "Emma",
		LastName:  "Thompson",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Liam",
		LastName:  "Chen",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Sophia",
		LastName:  "Rodriguez",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Ethan",
		LastName:  "Nguyen",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Olivia",
		LastName:  "Patel",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Noah",
		LastName:  "Kim",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Ava",
		LastName:  "Singh",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Jackson",
		LastName:  "O'Connor",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Isabella",
		LastName:  "Schmidt",
	},
	{
		ID:        uuid.NewString(),
		FirstName: "Mohammed",
		LastName:  "Al-Farsi",
	},
}
