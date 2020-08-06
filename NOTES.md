### usecase
POST /buy/ donde el payload es: {'id':1, 'cant':1} y la respuesta es: {'date':'31-07-2020'} => dt.now + 5

### orden llamados
main levanta servicios
server handlea http /buy/ => server llama a stock con grpc => stock llama a delivery con gprc => delivery calcula delivery_date y retorna delivery_date
server retorna json deliver_date

### exec
curl -i -X POST -d '{"id": 1, "quantity":1}' http://localhost:8080/buy/
res json '{"delivery_date":"random string delivery date"}'




## order to show code

1- protos
2- make file proto tool generator
3- protos.go generated from 2
4- how do we call our services? 1- create the grpc channel 2- importing our proto and using the `newXXXXXClient` 

main.go 
```go
case "server":
		// To call service methods, we first need to create a gRPC channel to communicate with the service.
		// We create the channel by passing the service address and port number to grpc.Dial()
		srv = services.NewServer(
			tracer,
			initGRPCConn(*stockaddr, tracer),
		)
```

services/server.go
```go
func NewServer(t opentracing.Tracer, stockconn *grpc.ClientConn) *Server {
	return &Server{
		stockClient: stock.NewStockClient(stockconn),
		tracer:      t,
	}
}
```

5- Sequence of events
The client calls the client stub. The call is a local procedure call, with parameters pushed on to the stack in the normal way.
The client stub packs the parameters into a message and makes a system call to send the message. Packing the parameters is called marshalling.
The client's local operating system sends the message from the client machine to the server machine.
The local operating system on the server machine passes the incoming packets to the server stub.
The server stub unpacks the parameters from the message. Unpacking the parameters is called unmarshalling.
Finally, the server stub calls the server procedure. The reply traces the same steps in the reverse direction.

