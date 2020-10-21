package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const LIM = 41

var facts [LIM]uint64

type Error struct {
	Msg string `json:msg`
}

var s Error

func main() {
	app := fiber.New()
	app.Get("/api/v1/fatorial/:number", Fatorial)
	app.Listen(":8080")
}

func Fatorial(c *fiber.Ctx) error {
	number := c.Params("number")
	if len(number) <= 0 {
		return c.Status(400).SendString(`{"msg":"error number obrigatorio!"}`)
	}

	result, err := Fact(number)
	if err != nil {
		s.Msg = err.Error()
		return c.Status(500).JSON(s)
	}
	return c.Status(200).SendString(result)
}

func Fact(number string) (res string, err error) {
	n, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return
	}

	result := factorial(n)
	res = strconv.FormatUint(result, 10)
	return
}

func factorial(n uint64) (res uint64) {

	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * factorial(n-1)
		return res
	}

	return 1
}
