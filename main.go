package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/muchobien/webhook-proxy/proxy"
)

func ProxyWebHooks(c *fiber.Ctx) error {
	service := c.Query("service")

	_, err := url.ParseRequestURI(service)
	if err != nil {
		return c.SendStatus(400)
	}

	if err := proxy.Do(c, service); err != nil {
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

	app.Post("/wh", ProxyWebHooks)

	app.Listen(fmt.Sprintf(":%s", port))
}
