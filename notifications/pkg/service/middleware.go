package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(NotificationsService) NotificationsService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificationsService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificationsService) NotificationsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "email", email, "content", content, "e0", e0)
	}()
	return l.next.SendEmail(ctx, email, content)
}