package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/sourcegraph/conc/pool"
	"github.com/spf13/cobra"
)

// concCmd represents the conc command
var concCmd = &cobra.Command{
	Use:   "conc",
	Short: "Conc command",
	Run: func(cmd *cobra.Command, args []string) {
		p := pool.New().WithMaxGoroutines(runtime.NumCPU())
		defer p.Wait()
		p.Go(func() {
			fmt.Println(time.Now())
		})
	},
}

func init() {
	rootCmd.AddCommand(concCmd)
}
