package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/iliyankg/colab-shield/backend/grpcserver"
)

func main() {
	viper.BindEnv("COLABSHIELD_PORT")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")
	viper.BindEnv("REDIS_PASSWORD")

	log.Info().Msg("Starting server...")

	port := viper.GetInt("COLABSHIELD_PORT")
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetInt("REDIS_PORT")
	redisPassword := viper.GetString("REDIS_PASSWORD")

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	// Not using the returned server for now
	_, err := grpcserver.Serve(port, redisClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start gRPC server")
	}
}
