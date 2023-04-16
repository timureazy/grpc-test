package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-demo/invoicer"
	"log"
	"net"
)

type invoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s invoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &invoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to server :%s", err)
	}
}
