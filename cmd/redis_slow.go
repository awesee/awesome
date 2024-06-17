package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/awesee/awesome/util"
	"github.com/spf13/cobra"
)

var redisSlowNumber int64

// redisSlowCmd represents the redisSlow command
var redisSlowCmd = &cobra.Command{
	Use:   "slow",
	Short: " Get top slow log",
	Run: func(cmd *cobra.Command, args []string) {
		rdb := redisClient()
		items, err := rdb.SlowLogGet(ctx, redisSlowNumber).Result()
		util.CheckErr(err)
		for k, v := range items {
			item, err := json.MarshalIndent(v, "", "\t")
			util.CheckErr(err)
			fmt.Printf("%d\t%s\n%s\n", k+1, v.Time, item)
		}
	},
}

func init() {
	redisCmd.AddCommand(redisSlowCmd)
	redisSlowCmd.Flags().Int64Var(&redisSlowNumber, "num", 10, "redis database number")
}
