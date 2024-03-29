package main

import (
	"github.com/iliyankg/colab-shield/cli/cmd"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cmd.Execute()
}
