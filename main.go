package main

import (
	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
	"github.com/sirius1024/overseer/config"
	"github.com/sirius1024/overseer/sniff"
)

func main() {
	conf := config.GetConfig()

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// health check url
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.POST("/ping", sniff.Pong)

	// check health right now
	if conf.Overseer.Self.APIEnabled {
		r.GET("/sniff", func(c *gin.Context) {
			sniff.Probe()
			sniff.IO()
			sniff.SelfReport()
			c.String(200, "done")
		})
	}

	// cronjob
	cronjob := cron.New(cron.WithSeconds())
	cronjob.AddFunc("* * * * * *", sniff.Probe)
	cronjob.AddFunc("* * * * * *", sniff.IO)
	cronjob.AddFunc(conf.Overseer.Self.Interval, sniff.SelfReport)
	cronjob.Start()

	r.Run(":" + string(conf.Port))
}
