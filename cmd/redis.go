package cmd

import (
	"fmt"
	"time"

	"github.com/awesee/awesome/util"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

var (
	redisHost   string
	redisPort   uint
	redisUser   string
	redisPass   string
	redisDb     int
	redisExpire time.Duration
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Lookup redis keys without expire",
	Run: func(cmd *cobra.Command, args []string) {
		rdb := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
			Username: redisUser,
			Password: redisPass,
			DB:       redisDb,
		})
		iter := rdb.Scan(ctx, 0, "*", 200).Iterator()
		util.CheckErr(iter.Err())
		idx := 0
		for iter.Next(ctx) {
			idx++
			ttl := rdb.TTL(ctx, iter.Val()).Val()
			verbosePrintf("%d\t%s\t%s\n", idx, iter.Val(), ttl)
			if verbose || ttl != -1 {
				continue
			}
			fmt.Println(iter.Val())
			if redisExpire > 0 {
				rdb.Expire(ctx, iter.Val(), redisExpire)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
	redisCmd.Flags().StringVar(&redisHost, "host", "127.0.0.1", "redis host")
	redisCmd.Flags().UintVar(&redisPort, "port", 6379, "redis port")
	redisCmd.Flags().StringVar(&redisUser, "user", "", "redis user")
	redisCmd.Flags().StringVar(&redisPass, "pass", "", "redis password")
	redisCmd.Flags().IntVar(&redisDb, "db", 0, "redis database number")
	redisCmd.Flags().DurationVar(&redisExpire, "expire", 0, `redis keys expire time, such as "1h", "20m" or "60s"`)
}
