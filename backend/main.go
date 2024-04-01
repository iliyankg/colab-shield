package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/iliyankg/colab-shield/backend/server"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 8080, "port to listen on")
var redisAddress = flag.String("redis-address", "localhost:6379", "address of the redis server")

func main() {
	flag.Parse()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Starting server...")

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     *redisAddress,
		Password: "",
		DB:       0,
	})

	// Create gRPC server
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterColabShieldServer(grpcServer, server.NewColabShieldServer(redisClient))

	// Listen on port
	log.Info().Msgf("Listening on port: %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	// Serve gRPC server
	log.Info().Msg("Serving gRPC")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
