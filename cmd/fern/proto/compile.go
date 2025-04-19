package proto

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	compileCmd = &cobra.Command{
		Use:   "compile",
		Short: "Compile proto files using Buf",
		Long: `Compile all proto files in the project using Buf.
	This requires buf to be installed (https://buf.build/docs/installation).`,
		Run: func(cmd *cobra.Command, args []string) {
			generateCmd := exec.Command("buf", "generate")
			if err := generateCmd.Run(); err != nil {
				// log.Fatal(err)
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		},
	}
)

func init() {
	ProtoCmd.AddCommand(compileCmd)
}
