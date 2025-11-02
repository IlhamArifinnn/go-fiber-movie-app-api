package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)
		Method := c.Method()
		Path := c.Path()
		Status := c.Response().StatusCode()
		IP := c.IP()

		log.Info().
			Str("method", Method).
			Str("path", Path).
			Int("status", Status).
			Dur("duration", duration).
			Str("ip", IP).
			Msg("Request completed")

		if err != nil {
			log.Error().
				Str("method", Method).
				Str("path", Path).
				Int("status", Status).
				Dur("duration", duration).
				Str("ip", IP).
				Err(err).
				Msg("Request failed")
		}
		return err
	}
}
