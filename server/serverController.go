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
	"Retail/workflow/status"
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

func (s *server) SaveUpdatePriceForApproval(ctx context.Context, productsPrice *workflow.ProductsRequest) (*workflow.ProductsResponse, error) {
	products := productsPrice.GetProducts()
	err := database.SavePriceInUpdateApprovalTable(products)
	return &workflow.ProductsResponse{Message: "successfully uploaded data"}, err
}

func (s *server) GetRecordsPendingForApproval(ctx context.Context, productsPrice *workflow.GetProductsRequest) (*workflow.GetProductsResponse, error) {

	records, err := database.GetAllPendingRecords();
	response := &workflow.GetProductsResponse{}
	if (err != nil) {
		log.Printf("Query failed while fetching pending entries for approval \n err: %v", err)
	} else {
		for records.Next() {
			var product_id int32;
			var version string;
			records.Scan(&product_id, &version)
			record := workflow.Product{
				ProductId: product_id,
				Version: version}
			response.Products = append(response.Products, &record)
		}
	}
	return response, err
}

func (s *server) UpdateStatusToCompleted(ctx context.Context, request *workflow.ProductsRequest) (*workflow.ProductsResponse, error) {
	records := request.GetProducts();
	err := database.ChangeStatusTo(status.COMPLETED, records);

	message := fmt.Sprintf("Successfully changed status of %v to picked", records);
	if (err != nil) {
		message = fmt.Sprintf("Failed to change status of %v to picked", records);
	}
	return &workflow.ProductsResponse{Message: message}, err
}