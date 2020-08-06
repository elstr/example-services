package services

import (
	"encoding/json"
	"fmt"
	stock "github.com/elstr/example-services/proto/stock"
	"github.com/elstr/example-services/trace"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

// TODO: Better naming.. server is :S

// Server struct that contains the stock service client to be called
type Server struct {
	stockClient stock.StockClient
	tracer      opentracing.Tracer
}

// Item represents a product with a stock
type Item struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}

// NewServer returns a new server
func NewServer(t opentracing.Tracer, stockconn *grpc.ClientConn) *Server {
	return &Server{
		stockClient: stock.NewStockClient(stockconn),
		tracer:      t,
	}
}

// Run the server
func (s *Server) Run(port int) error {
	mux := trace.NewServeMux(s.tracer)
	mux.Handle("/buy", http.HandlerFunc(s.buyHandler))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func (s *Server) buyHandler(w http.ResponseWriter, r *http.Request) {
	reqData := Item{}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("buyHandler - req data: %+v", &reqData)
	ctx := r.Context()
	res, err := s.stockClient.UpdateStock(ctx, &stock.Request{
		Item:     reqData.ID,
		Quantity: reqData.Quantity,
	})

	log.Println(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
