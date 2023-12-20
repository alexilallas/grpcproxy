package main

import (
	"log"
	"net"
	"net/http"

	"github.com/alexilallas/grpcproxy/internal/adapter/inbound"
	"github.com/alexilallas/grpcproxy/internal/adapter/outbound/gateway"
	"github.com/alexilallas/grpcproxy/internal/config"
	"github.com/alexilallas/grpcproxy/internal/core/usecase"
	pb "github.com/alexilallas/grpcproxy/internal/grpc"
	"google.golang.org/grpc"
)

const serverPort = ":8080"

func main() {
	list, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalln(err)
	}

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	taskUsecase := usecase.ProvideTaskUseCase(gateway.ProvideTaskGateway(http.Client{}, conf.GetTaskURL(), conf.GetTaskToken()))
	grpcServer := grpc.NewServer()
	pb.RegisterTaskManagerServer(grpcServer, inbound.ProvideNewHandler(taskUsecase))

	log.Println("listing on port ", serverPort)
	if err = grpcServer.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
