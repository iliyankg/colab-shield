package client

import (
	"context"

	pb "github.com/iliyankg/colab-shield/protos"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewColabShieldClient(ctx context.Context, addr string) (*grpc.ClientConn, pb.ColabShieldClient) {
	// FIXME: Fix the insecure.NewCredentials() call
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect")
	}

	return conn, pb.NewColabShieldClient(conn)
}
