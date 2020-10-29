package main

import (
	"fmt"

	lb "github.com/dayvson/go-leaderboard"
	"github.com/ghiffaryuthian/leaderboard-sample/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	leaderboard := lb.NewLeaderboard(lb.RedisSettings{
		Host:     cfg.GetString("REDIS_ADDR"),
		Password: cfg.GetString("REDIS_PASSWORD"),
	}, "leaderboard", 10)

	if _, err := leaderboard.RankMember("pepega", 69); err != nil {
		panic(err)
	}

	fmt.Println(leaderboard.TotalMembers())

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run()
}
