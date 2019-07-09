package service

import (
	"context"
	"kit-test/notifications/pkg/grpc/pb"
	"log"

	"github.com/go-kit/kit/sd/etcd"

	"google.golang.org/grpc"
)

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, email string) error
}

type basicUsersService struct {
	notificatorServiceClient pb.NotificationsClient
}

func (b *basicUsersService) Create(ctx context.Context, email string) error {
	// TODO implement the business logic of Create
	reply, err := b.notificatorServiceClient.SendEmail(
		context.Background(),
		&pb.SendEmailRequest{Email: email, Content: "New user created!"},
	)
	if reply != nil {
		log.Printf("Email ID: %s\n", reply.Id)
	}
	return err
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {

	var (
		etcdServer = "http://etcd:2379"
		prefix     = "/services/notifications/"
	)

	client, err := etcd.NewClient(context.Background(), []string{etcdServer}, etcd.ClientOptions{})
	if err != nil {
		log.Printf("Unable to connect to etcd: %s\n", err.Error())
	}

	entries, err := client.GetEntries(prefix)
	if err != nil || len(entries) == 0 {
		log.Printf("Unable to get entries: %s\n", err.Error())
		return new(basicUsersService)
	}

	conn, err := grpc.Dial(entries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("Unable to connect to notificator: %s\n", err.Error())
	} else {
		log.Printf("Found and connected to notifications service: %s", entries[0])
	}
	return &basicUsersService{
		notificatorServiceClient: pb.NewNotificationsClient(conn),
	}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
