// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "kit-test/notifications/pkg/endpoint"
	pb "kit-test/notifications/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	sendEmail grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.NotificationsServer {
	return &grpcServer{sendEmail: makeSendEmailHandler(endpoints, options["SendEmail"])}
}
