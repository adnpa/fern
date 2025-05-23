package main

import (
	"log"

	"github.com/adnpa/fern/cmd/fern/proto"
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
