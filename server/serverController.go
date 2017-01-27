package server

import (
	"net"
	"log"
	"google.golang.org/grpc"
	workflow "github.com/RetailMarket/workFlowClient"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	"Retail/workflow/database"
	"fmt"
)

const (
	port = ":4000"
)

type server struct{}

func CreateServerConnection() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	_server := grpc.NewServer()

	workflow.RegisterWorkFlowServer(_server, &server{})

	reflection.Register(_server)
	fmt.Printf("Listening to port %s\n", port);
	if err := _server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SaveUpdatePriceForApproval(ctx context.Context, productsPrice *workflow.PriceUpdateRequest) (*workflow.PriceUpdateResponse, error) {
	products := productsPrice.GetProducts()
	err := database.SavePriceInUpdateApprovalTable(products)
	return &workflow.PriceUpdateResponse{Message: "successfully uploaded data"}, err
}