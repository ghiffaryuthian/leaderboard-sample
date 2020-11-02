package main

import (
	"context"
	"fmt"

	"github.com/ghiffaryuthian/leaderboard-sample/config"
	"github.com/ghiffaryuthian/leaderboard-sample/handler"
	"github.com/ghiffaryuthian/leaderboard-sample/leaderboard"
	redis "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.MakeLeaderboardHandler(e, lbRepo)
	handler.MakeMiscHandler(e)

	e.Logger.Fatal(e.Start(":1234"))

}
