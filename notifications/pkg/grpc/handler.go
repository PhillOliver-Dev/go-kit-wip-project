package grpc

import (
	"context"
	endpoint "kit-test/notifications/pkg/endpoint"
	pb "kit-test/notifications/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
func decodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SendEmailRequest)
	return endpoint.SendEmailRequest{Email: req.Email, Content: req.Content}, nil
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.SendEmailResponse)
	return &pb.SendEmailReply{Id: reply.Id}, nil
}

func (g *grpcServer) SendEmail(ctx context1.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	_, rep, err := g.sendEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailReply), nil
}
