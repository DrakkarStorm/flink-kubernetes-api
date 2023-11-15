package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Créer une instance de l'application Fiber
	app := fiber.New()

	// Définir une route pour le chemin "/api/hello"
	app.Get("/api/deployments", func(c *fiber.Ctx) error {
		deployments, err := ListFlinkDeployment()
		if err != nil {
			// Renvoyer le message au format JSON
			return c.Status(500).SendString(err.Error())
		}

		// Renvoyer le message au format JSON
		return c.JSON(deployments)
	})

	// Définir une route pour le chemin "/api/hello"
	app.Get("/api/jobs", func(c *fiber.Ctx) error {
		jobs, err := ListFlinkJobs()
		if err != nil {
			// Renvoyer le message au format JSON
			return c.Status(500).SendString(err.Error())
		}

		// Renvoyer le message au format JSON
		return c.JSON(jobs)
	})

	app.Patch("/api/deployments/:name", func(c *fiber.Ctx) error {
		deployments, err := ListFlinkDeployment()
		if err != nil {
			// Renvoyer le message au format JSON
			return c.Status(500).SendString(err.Error())
		}

		// Renvoyer le message au format JSON
		return c.JSON(deployments)
	})

	app.Patch("/api/jobs/:id", func(c *fiber.Ctx) error {
		deployments, err := ListFlinkDeployment()
		if err != nil {
			// Renvoyer le message au format JSON
			return c.Status(500).SendString(err.Error())
		}

		// Renvoyer le message au format JSON
		return c.JSON(deployments)
	})

	// Démarrer le serveur sur le port 8080
	app.Listen(":8080")
}
