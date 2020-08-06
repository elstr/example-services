# gRPC example 

### Usecase
POST `/buy` where payload is: `{'id':1, 'cant':1}` and the res is: `{'date':'31-07-2020'}` (fake delivery date)

#### Order of calls
main starts all the services <br />
server handles http reqs to `/buy` <br />
server calls stock service with a rpc call <br />
stock service callas delivery service with a rpc call <br />
delivery service retuns fake delivery date to stock and stock sends date to server

### exec
```
curl -i -X POST -d '{"id": 1, "quantity":1}' http://localhost:8080/buy/ 
```

`res json '{"delivery_date":"random string delivery date"}'`
