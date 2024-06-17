package cmd

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

const VERSION = "0.1.1"

var (
	ctx     = context.Background()
	verbose = false
)
var exeSuffix = func() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "awesome",
	Short:            "Awesome tools",
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func verbosePrintln(a ...any) {
	if verbose {
		fmt.Println(a...)
	}
}

func verbosePrintf(format string, a ...any) {
	if verbose {
		fmt.Printf(format, a...)
	}
}
