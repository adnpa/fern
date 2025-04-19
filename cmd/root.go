package cmd

import (
	"fmt"
	"os"

	"github.com/adnpa/fern/cmd/proto"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fern",
	Short: "fern: An elegant toolkit for Go microservices.",
	Long:  `fern: An elegant toolkit for Go microservices.`,
}

func init() {
	rootCmd.AddCommand(proto.ProtoCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
