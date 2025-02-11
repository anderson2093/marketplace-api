// cmd/api/main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2" // Puedes usar Fiber, Gin o el framework que prefieras
	// Importa otros paquetes internos según lo necesites, por ejemplo, tu configuración o handlers.
)

func main() {
	app := fiber.New()

	// Configuración de middlewares (CORS, logging, etc.) aquí
	// app.Use(...)

	// Configuración de rutas o endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido al Marketplace de Consolidación de Carga")
	})

	// Aquí puedes registrar otros endpoints desde tus handlers
	// por ejemplo: userHandler.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
