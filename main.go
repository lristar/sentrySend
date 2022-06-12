package main

import (
	"fmt"
	wxwork "monitorD/wx_work"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://3f9943cbdcf74303b0d600b1c8a521b4@192.168.195.223:9000/2",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					// You have access to the original Request here
					fmt.Println(req)
				}
			}
			return event
		},
	})
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	app := gin.Default()

	app.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	app.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	})

	app.GET("/", func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
				hub.CaptureMessage("User provided unwanted query string, but we recovered just fine")
			})
		}
		ctx.Status(http.StatusOK)
	})

	app.GET("/foo", func(ctx *gin.Context) {
		// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
		// in the middleware before, it will be sent through as well
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "错误的请求"})
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.CaptureMessage("错误的请求")
		}
	})
	app.POST("/sentry/callback", func(ctx *gin.Context) {
		m := make(map[string]interface{}, 0)
		ctx.ShouldBind(&m)
		wxwork.Send(m["message"].(string))
	})
	app.Run(":3000")
}
