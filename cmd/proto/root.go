package proto

import "github.com/spf13/cobra"

var (
	ProtoCmd = &cobra.Command{
		Use:   "proto",
		Short: "protobuf generate and compile",
	}
)
