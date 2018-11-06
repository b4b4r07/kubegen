package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// Version is the version string
const Version = "0.1.0"

var (
	showVersion bool
)

// RootCmd is the entry point
var RootCmd = &cobra.Command{
	Use:           "kubegen",
	Short:         "Kubernetes manifest generator",
	Long:          "kubegen - Kubernetes manifest generator",
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("version %s/%s\n", Version, runtime.Version())
			return
		}
		cmd.Usage()
	},
}

// Execute is the entry function
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	initConf()
	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show the version and exit")
}

func initConf() {
}
