package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/zdacoder/go-fiber-movie-app-api/internal/handlers"
	"github.com/zdacoder/go-fiber-movie-app-api/internal/middlewares"
)

func Init(app *fiber.App) {
	// Apply middlewares
	app.Use(middlewares.LoggerMiddleware())

	// CORS middleware
	app.Use(middlewares.CORSMiddleware())

	// Movie routes
	movies := app.Group("/api/movies")
	movies.Get("/", handlers.ListMovies)
	movies.Get("/:id", handlers.GetMovie)
	movies.Post("/", handlers.CreateMovie)
	movies.Put("/:id", handlers.UpdateMovie)
	movies.Delete("/:id", handlers.DeleteMovie)

	// Swagger documentation route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// 404 Handler for undefined routes
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Resource not found")
	})
}
