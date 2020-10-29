package main

import (
	"fmt"

	"github.com/ghiffaryuthian/leaderboard-sample/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg.GetString("redis.host") + ":" + cfg.GetString("redis.port") + " | " + viper.GetString("redis.password"+"|"))
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run()
}
