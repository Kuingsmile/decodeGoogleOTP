package cmd

import (
	"os"
	"runtime"

	"github.com/kuingsmile/decodeGoogleOTP/conf"
	"github.com/spf13/cobra"
)

var commandVersion = &cobra.Command{
	Use:   "version",
	Short: "Print current version of the application",
	Run:   printVersion,
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(commandVersion)
}

func printVersion(cmd *cobra.Command, args []string) {
	version := "decodeGoogleOTP\nversion " + conf.Version + "\n"
	version += "OS: " + runtime.GOOS + "\n"
	version += "Architecture: " + runtime.GOARCH + "\n"
	version += "Go Version: " + runtime.Version() + "\n"

	os.Stdout.WriteString(version)
}
