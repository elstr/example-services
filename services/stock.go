package services

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	stock "github.com/grpc-example/example-services/proto"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Stock Server implements the stock service
type Stock struct {
	Items  map[int]int
	tracer opentracing.Tracer
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
