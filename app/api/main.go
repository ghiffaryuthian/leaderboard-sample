package main

import (
	"context"
	"fmt"

	"github.com/ghiffaryuthian/leaderboard-sample/config"
	"github.com/ghiffaryuthian/leaderboard-sample/leaderboard"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
)

func main() {
	cfg := config.NewConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.GetString("REDIS_ADDR"),
		Password: cfg.GetString("REDIS_PASSWORD"),
		DB:       cfg.GetInt("REDIS_DB"),
	})

	leaderboardRepo := leaderboard.NewRedisRepo(rdb, "leaderboard", 10)

	if _, err := leaderboardRepo.RankMember(context.TODO(), "andi", 8869); err != nil {
		panic(err)
	}

	fmt.Println(leaderboardRepo.TotalMembers(context.TODO()))

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run()
}
