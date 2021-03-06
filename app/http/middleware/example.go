package middleware

import (
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/request"
)

func Example() request.HandlerFunc {
	return func(c request.Context) {
		t := zone.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := zone.Since(t)
		log.Info("latency", tmaic.V{"latency": latency})

		// access the status we are sending
		status := c.Writer.Status()
		log.Info("status", tmaic.V{"status": status})
	}
}
