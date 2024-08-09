package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	First  float64 `json:"first"`
	Second float64 `json:"second"`
}

func main() {
	app := fiber.New()

	v1 := app.Group("api/v1")

	v1.Get("/add", func(c *fiber.Ctx) error {
		p := new(Request)
		if err := c.BodyParser(p); err != nil {
			log.Println(err)
			return fiber.ErrBadRequest
		}

		return c.JSON(fiber.Map{
			"result": p.First + p.Second,
		})
	})

	v1.Get("/subtract", func(c *fiber.Ctx) error {
		p := new(Request)
		if err := c.BodyParser(p); err != nil {
			log.Println(err)
			return fiber.ErrBadRequest
		}

		return c.JSON(fiber.Map{
			"result": p.First - p.Second,
		})
	})

	v1.Get("/multiply", func(c *fiber.Ctx) error {
		p := new(Request)
		if err := c.BodyParser(p); err != nil {
			log.Println(err)
			return fiber.ErrBadRequest
		}

		return c.JSON(fiber.Map{
			"result": p.First * p.Second,
		})
	})

	v1.Get("/divide", func(c *fiber.Ctx) error {
		p := new(Request)
		if err := c.BodyParser(p); err != nil {
			log.Println(err)
			return fiber.ErrBadRequest
		}
		if p.Second == 0 {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "Cannot divide by zero",
			})
		}

		return c.JSON(fiber.Map{
			"result": p.First / p.Second,
		})
	})

	log.Fatal(app.Listen(":4000"))
}
