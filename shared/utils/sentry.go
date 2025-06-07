package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/gin" 
	"github.com/gin-gonic/gin"
)

func InitSentry(serviceName string) {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		log.Fatal("SENTRY_DSN not set")
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		EnableTracing:    true,
		TracesSampleRate: 1.0, 
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			event.Tags = map[string]string{
				"service": serviceName,
			}
			return event
		},
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}
	fmt.Printf("Sentry initialized for service: %s\n", serviceName)
}

func SentryMiddleware() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{})
}
