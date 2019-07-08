package service

import "context"

// NotificationsService describes the service.
type NotificationsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmail(ctx context.Context, email string, content string) error
}
