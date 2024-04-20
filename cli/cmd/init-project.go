package cmd

import (
	"context"
	"time"

	"github.com/iliyankg/colab-shield/cli/client"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var initProjectCmd = &cobra.Command{
	Use:   "init-project",
	Short: "Initializes a project on the backend.",
	Long:  `Initializes a project on the backend.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		conn, client := client.NewColabShieldClient(serverAddress)
		defer conn.Close()

		payload := &protos.InitProjectRequest{
			ProjectId: gitBranch,
		}

		response, err := client.InitProject(ctx, payload)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to init project")
		}

		if response.Status != protos.Status_OK {
			log.Fatal().Msg("status not OK")
		}
	},
}
