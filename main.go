package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/muchobien/webhook-proxy/proxy"
)

const (
	// ServicesURI is the URI for the services
	ServicesURI string = "https://gist.githubusercontent.com/KROSF/22dc413f61212be0677efbf04f3a1c66/raw/services.json"
)

func ProxyWebHooks(c *fiber.Ctx) error {
	agent := fiber.AcquireClient()

	services := make(map[string]string)
	agent.Get(ServicesURI).Struct(&services)
	service, ok := services[c.Params("service")]

	if !ok {
		return c.SendStatus(404)
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

	app.Post("/wh/:service", ProxyWebHooks)

	app.Listen(fmt.Sprintf(":%s", port))
}
