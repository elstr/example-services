package services

import (
	"fmt"
	"net"

	delivery "github.com/elstr/example-services/proto/delivery"
	stock "github.com/elstr/example-services/proto/stock"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Stock Server implements the stock service
type Stock struct {
	Items          map[int]int
	deliveryClient delivery.DeliveryClient
	tracer         opentracing.Tracer
}

// Run starts the server
func (s *Stock) Run(port int) error {
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer),
		),
	)

	stock.RegisterStockServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

// NewStock returns a new server
func NewStock(t opentracing.Tracer, deliveryConn *grpc.ClientConn) *Stock {
	return &Stock{
		Items: map[int]int{
			1: 5,
			2: 10,
		},
		deliveryClient: delivery.NewDeliveryClient(deliveryConn),
		tracer:         t,
	}
}

// UpdateStock will update quantity for a given item and call delivery service to get the delivery date
func (s *Stock) UpdateStock(ctx context.Context, req *stock.Request) (*stock.Response, error) {
	// 1. get request quantity
	// 2. discount stock
	// 3. call delivery service and get delivery date
	// 4. return delivery date as stock.Response
	return nil, nil
}
