package database

import "github.com/rs/zerolog/log"

func Migrate(models ...interface{}) {
	if DB == nil {
		log.Fatal().Msg("Database connection is not established")
	}

	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatal().Err(err).Msg("Database migration failed")
	}

	log.Info().Msg("Database migration completed successfully")
}
