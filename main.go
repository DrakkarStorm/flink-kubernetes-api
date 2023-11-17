package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Status struct {
	State string `json:"state"`
}

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
		log.Println("Received a GET request")
		jobs, err := ListFlinkJobs()
		if err != nil {
			// Renvoyer le message au format JSON
			return c.Status(500).SendString(err.Error())
		}

		// Renvoyer le message au format JSON
		return c.JSON(jobs)
	})

	app.Patch("/api/jobs/:name", func(c *fiber.Ctx) error {
		jobName := c.Params("name")

		// Create an instance of the User struct to store the parsed JSON
		status := new(Status)
		// Parse the JSON body into the User struct
		if err := c.BodyParser(status); err != nil {
			return err
		}

		if err := UpdateFlinkSessionJob(jobName, status.State); err != nil {
			// Renvoyer le message au format JSON
			return c.Status(500).SendString(err.Error())
		}

		// Renvoyer le message au format JSON
		return c.Status(201).Send(nil)
	})

	// Démarrer le serveur sur le port 8080
	app.Listen(":8080")
}
