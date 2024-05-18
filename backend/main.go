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
	"github.com/iliyankg/colab-shield/backend/redisdatabase"
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

	colabDatabase := redisdatabase.NewRedisDatabase(redisClient)

	group, _ := errgroup.WithContext(context.Background())

	grpcServer := grpcserver.NewColabShieldServer(colabDatabase)
	httpServer := httpserver.NewColabShieldServer(colabDatabase)

	group.Go(func() error {
		_, err := grpcServer.Serve(grpcPort)
		return err
	})
	group.Go(func() error {
		return httpServer.Serve(httpPort)
	})

	if err := group.Wait(); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
