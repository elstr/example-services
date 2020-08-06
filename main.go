package main

import (
	"flag"
	"fmt"
	"github.com/elstr/example-services/dialer"
	"github.com/elstr/example-services/services"
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
	case "server":
		// To call service methods, we first need to create a gRPC channel to communicate with the service.
		// We create the channel by passing the service address and port number to grpc.Dial()
		srv = services.NewServer(
			tracer,
			initGRPCConn(*stockaddr, tracer),
		)
	case "delivery":
		srv = services.NewDelivery(tracer)
	case "stock":
		srv = services.NewStock(
			tracer,
			initGRPCConn(*deliveryaddr, tracer),
		)

	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}

	log.Println("Running - " + os.Args[1])
	srv.Run(*port)
}

func initGRPCConn(addr string, tracer opentracing.Tracer) *grpc.ClientConn {
	// Dial receives the address + dial options (TLS, GCE credentials, JWT credentials, tracer, etc)
	conn, err := dialer.Dial(addr, dialer.WithTracer(tracer))
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}
	return conn
}
