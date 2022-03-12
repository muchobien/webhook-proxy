package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
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

	req := agent.Post(service).Body(c.Body())

	headers := c.GetReqHeaders()
	for k, v := range headers {
		if strings.ToLower(k) != "connection" {
			req.Set(k, v)
		}
	}

	code, _, _ := req.Bytes()
	if code > 399 {
		return c.SendStatus(code)
	}

	return c.SendString(service)
}

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Webhook Proxy!")
	})

	app.Post("/webhook/:service", ProxyWebHooks)

	app.Listen(fmt.Sprintf(":%s", port))
}
