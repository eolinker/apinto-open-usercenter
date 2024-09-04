package main

import (
	grpc_service "github.com/eolinker/apinto-dashboard/grpc-service"
	"github.com/eolinker/eosc/common/bean"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpc(addr string) {
	grpcConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := grpc_service.NewGetConsoleInfoClient(grpcConn)
	bean.Injection(&client)
}
