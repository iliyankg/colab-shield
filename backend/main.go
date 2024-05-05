package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"

	"github.com/iliyankg/colab-shield/backend/grpcserver"
	"github.com/iliyankg/colab-shield/backend/httpserver"
)

func main() {
	viper.BindEnv("COLABSHIELD_GRPC_PORT")
	viper.BindEnv("COLABSHIELD_HTTP_PORT")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")
	viper.BindEnv("REDIS_PASSWORD")

	log.Info().Msg("Starting server...")

	grpcPort := viper.GetInt("COLABSHIELD_GRPC_PORT")
	httpPort := viper.GetInt("COLABSHIELD_HTTP_PORT")
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetInt("REDIS_PORT")
	redisPassword := viper.GetString("REDIS_PASSWORD")

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		_, err := grpcserver.Serve(grpcPort, redisClient)
		return err
	})
	group.Go(func() error {
		return httpserver.Serve(httpPort, redisClient)
	})

	if err := group.Wait(); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
