module github.com/iliyankg/colab-shield/cli

go 1.22.0

require (
	github.com/rs/zerolog v1.32.0
	github.com/spf13/cobra v1.8.0
	github.com/iliyankg/colab-shield/protos v0.0.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/iliyankg/colab-shield/protos v0.0.0 => ../protos