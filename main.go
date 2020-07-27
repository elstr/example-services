package main

import (
	"flag"
	"fmt"
	services "github.com/elstr/example-services"
	"github.com/elstr/example-services/dialer"
	"github.com/elstr/example-services/trace"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"log"
	"os"
)

type server interface {
	Run(int) error
}

func main() {
	var (
		port         = flag.Int("port", 8080, "The service port")
		jaegeraddr   = flag.String("jaeger", "jaeger:6831", "Jaeger address")
		stockaddr    = flag.String("stockaddr", "stock:8080", "Stock service addr")
		deliveryaddr = flag.String("deliveryaddr", "delivery:8080", "Delivery server addr")
	)

	flag.Parse()

	tracer, err := trace.New("stock", *jaegeraddr)
	if err != nil {
		log.Fatalf("trace new error: %v", err)
	}

	var srv server

	// this will be called with docker-compose
	switch os.Args[1] {
	case "delivery":
		srv = services.NewDelivery(tracer)
	case "stock":
		srv = services.NewStock(
			tracer,
			initGRPCConn(*deliveryaddr, tracer),
		)
	case "server":
		srv = services.NewServer(
			tracer,
			initGRPCConn(*stockaddr, tracer),
		)
	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}

	srv.Run(*port)
}

func initGRPCConn(addr string, tracer opentracing.Tracer) *grpc.ClientConn {
	conn, err := dialer.Dial(addr, dialer.WithTracer(tracer))
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}
	return conn
}
