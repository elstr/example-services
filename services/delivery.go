package services

import (
	"fmt"
	delivery "github.com/elstr/example-services/proto/delivery"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

// DeliveryServer implements the Delivery service
type DeliveryServer struct {
	deliveryDate string
	tracer       opentracing.Tracer
}

// Run instaciates a new Delivery server
// Once we’ve implemented all our methods, we also need to start up a gRPC server so that clients can actually use our service.
func (s *DeliveryServer) Run(port int) error {
	// Create an instance of the gRPC server using
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer),
		),
	)

	// Register our service implementation with the gRPC server
	delivery.RegisterDeliveryServer(srv, s)

	// Specify the port we want to use to listen for client requests using
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called
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
	log.Println("Delivery - Quantity: " + strconv.Itoa(int(req.GetQuantity())))
	s.deliveryDate = "01/01/2021"

	return &delivery.Response{DeliveryDate: s.deliveryDate}, nil
}
