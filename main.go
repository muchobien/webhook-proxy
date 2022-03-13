package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ProxyWebHooks(c *fiber.Ctx) error {
	upstream := c.Query("upstream")

	_, err := url.ParseRequestURI(upstream)
	if err != nil {
		return c.SendStatus(400)
	}

	if err := proxy.Do(c, upstream); err != nil {
		return err
	}

	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Webhook Proxy!")
	})

	app.Post("/", ProxyWebHooks)

	app.Listen(fmt.Sprintf(":%s", port))
}
