protos :
	protoc -I proto/ stock.proto --go_out=plugins=grpc:proto;
	protoc -I proto/ delivery.proto --go_out=plugins=grpc:proto;

run:
	docker-compose build
	docker-compose up --remove-orphans