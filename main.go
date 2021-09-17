package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"http2/srv/doctor"
	"http2/srv/patient"
)

var rCMD = &cobra.Command{
	Use:     "grpc-test",
	Short:   "",
	Long:    "",
	Example: "",
}

var log *zap.Logger

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("unable to create logger: %w", err))
	}
	log = l.Named("main")

	rCMD.AddCommand(
		patient.CMD,
		doctor.CMD,
	)
}

func main() {
	// Listen to the kill command in the background to not cause issues whilst executing commands
	// go killable.ListenToKill()

	if err := rCMD.Execute(); err != nil {
		log.Fatal("failed to execute command", zap.Error(err))
		os.Exit(1)
	}
}
