package services

import (
	"fmt"
	"net"

	delivery "github.com/elstr/grpc-example/proto/delivery"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// DeliveryServer implements the Delivery service
type DeliveryServer struct {
	deliveryDate string
	tracer       opentracing.Tracer
}

// Run starts the server
func (s *DeliveryServer) Run(port int) error {
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer),
		),
	)

	delivery.RegisterDeliveryServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

// NewDelivery returns a new server
func NewDelivery(t opentracing.Tracer) *DeliveryServer {
	return &DeliveryServer{
		deliveryDate: "",
		tracer:       t,
	}
}

// GetDeliveryDate will return a random delivery date
func (s *DeliveryServer) GetDeliveryDate(ctx context.Context, req *delivery.Request) (*delivery.Response, error) {
	// 1. get request quantity
	// 2. generate random delivery date
	// 3. return delivery date as delivery.Response
	return nil, nil
}
