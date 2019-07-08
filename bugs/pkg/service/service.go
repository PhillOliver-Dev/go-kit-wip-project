package service

import "context"

// BugsService describes the service.
type BugsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	Create(ctx context.Context, s string) (rs string, err error)
}

type basicBugsService struct{}

func (b *basicBugsService) Create(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Create
	return rs, err
}

// NewBasicBugsService returns a naive, stateless implementation of BugsService.
func NewBasicBugsService() BugsService {
	return &basicBugsService{}
}

// New returns a BugsService with all of the expected middleware wired in.
func New(middleware []Middleware) BugsService {
	var svc BugsService = NewBasicBugsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
