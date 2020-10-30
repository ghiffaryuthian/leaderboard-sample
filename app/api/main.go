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

	lbRepo := leaderboard.NewRedisRepo(rdb, "leaderboard", 10)
	lbService := leaderboard.NewService(lbRepo)

	userDetail, err := lbService.RankMember(context.TODO(), "ayaya", 969696)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted %s rank:%d | score:%.0f\n", userDetail.Name, userDetail.Rank, userDetail.Score)

	rank, _ := lbRepo.GetUserRank(context.TODO(), "ayaya")
	score, _ := lbRepo.GetUserScore(context.TODO(), "ayaya")
	memberCount, _ := lbRepo.TotalMembers(context.TODO())
	fmt.Printf("ayaya rank:%d | score:%.0f\nmember count:%d", rank, score, memberCount)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run()
}
