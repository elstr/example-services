### usecase
POST /buy/ donde el payload es: {'id':1, 'cant':1} y la respuesta es: {'date':'31-07-2020'} => dt.now + 5

### orden llamados
main levanta servicios
server handlea http /buy/ => server llama a stock con grpc => stock llama a delivery con gprc => delivery calcula delivery_date y retorna delivery_date
server retorna json deliver_date

### exec
curl -i -X POST -d '{"id": 1, "quantity":1}' http://localhost:8080/buy/
res json '{"delivery_date":"random string delivery date"}'


