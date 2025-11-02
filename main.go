package main

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/zdacoder/go-fiber-movie-app-api/config"
	"github.com/zdacoder/go-fiber-movie-app-api/config/database"

	"github.com/zdacoder/go-fiber-movie-app-api/internal/models"
	"github.com/zdacoder/go-fiber-movie-app-api/internal/routes"
	"github.com/zdacoder/go-fiber-movie-app-api/internal/validators"
	"github.com/zdacoder/go-fiber-movie-app-api/pkg/logger"
)

func main() {
	// Load environment variables
	config := config.Load()

	// Initialize logger
	logger.Init(config)

	// Initialize database connection
	database.Connect(config)

	// Run database migrations
	database.Migrate(&models.Movie{})

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	// validation initialization
	validators.Init()

	// Initialize routes
	routes.Init(app)

	// Construct server address and start the server
	addr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)

	// Start the server
	log.Info().Msgf("Starting server on %s in %s mode", addr, config.AppEnv)

	// If server fails to start, log the error and exit
	if err := app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
