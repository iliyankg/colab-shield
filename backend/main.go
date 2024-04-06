package main

import (
	"fmt"
	"net"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/iliyankg/colab-shield/backend/server"
	pb "github.com/iliyankg/colab-shield/protos"
)

func main() {
	viper.BindEnv("COLABSHIELD_PORT")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")
	viper.BindEnv("REDIS_PASSWORD")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

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

	// Create gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(server.UnaryInterceptor),
	)
	pb.RegisterColabShieldServer(grpcServer, server.NewColabShieldServer(redisClient))

	// Listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Msgf("Listening on port: %d", port)

	// Serve gRPC server
	log.Info().Msg("Serving gRPC")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
