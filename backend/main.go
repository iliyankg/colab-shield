package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/om"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/iliyankg/colab-shield/backend/grpcserver"
	"github.com/iliyankg/colab-shield/backend/models"
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

	rueidisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{fmt.Sprintf("%s:%d", redisHost, redisPort)},
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start rueidis client")
	}

	// TODO: Repo per project maybe?
	repo := om.NewJSONRepository("file_infos", models.FileInfo{}, rueidisClient)
	if _, ok := repo.(*om.JSONRepository[models.FileInfo]); !ok {
		repo.CreateIndex(context.Background(), func(schema om.FtCreateSchema) rueidis.Completed {
			return schema.FieldName("$.fileId").As("fileId").Tag().Build()
		})
		repo.CreateIndex(context.Background(), func(schema om.FtCreateSchema) rueidis.Completed {
			return schema.FieldName("$.userIds").As("userIds").Tag().Build()
		})
		repo.CreateIndex(context.Background(), func(schema om.FtCreateSchema) rueidis.Completed {
			return schema.FieldName("$.userIds").As("userIds").Tag().Build()
		})
		repo.CreateIndex(context.Background(), func(schema om.FtCreateSchema) rueidis.Completed {
			return schema.FieldName("$.claimMode").As("claimMode").Tag().Build()
		})
	}

	repo.Search(context.Background(), func(search om.FtSearchIndex) rueidis.Completed {
		return search.Query("@fileId:($fi)").Build()
	})

}
