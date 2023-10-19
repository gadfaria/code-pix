# code-pix

* Docker
* Golang
* Apache Kafka
* Postgres

* Gerar o docker -> docker-compose up -d
* Verificar os containers -> docker-compose ps
* Entrar no container -> docker exec -it {NOME_CONTAINER} bash


# gRPC
RPC - Remote Procedure Call

## Protocol Buffers

## HTTP/2  (SPDY)


protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto