package httpserver

import (
	"fmt"
	"net/http"

	"github.com/alexliesenfeld/health"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func Serve(port int, redisClient *redis.Client) error {
	http.HandleFunc("/health", health.NewHandler(newHealthChecker(redisClient)))

	log.Info().Msgf("Http listening on port: %d", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		return err
	}

	return nil
}
