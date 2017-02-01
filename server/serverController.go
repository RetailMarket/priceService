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

func (s *server) PendingRecords(ctx context.Context, productsPrice *workflow.Request) (*workflow.Records, error) {

	records, err := database.GetAllPendingRecords();
	response := &workflow.Records{}
	if (err != nil) {
		log.Printf("Query failed while fetching pending entries for approval \n err: %v", err)
	} else {
		for records.Next() {
			var product_id int32;
			var version string;
			records.Scan(&product_id, &version)
			record := workflow.Entry{
				ProductId: product_id,
				Version: version}
			response.Entries = append(response.Entries, &record)
		}
	}
	return response, err
}

func (s *server) NotifyRecordsPicked(ctx context.Context, productsPrice *workflow.Records) (*workflow.Response, error) {
	products := productsPrice.GetEntries()
	err := database.SavePriceInUpdateApprovalTable(products)
	return &workflow.Response{Message: "successfully uploaded data"}, err
}

func (s *server) NotifyRecordsProcessed(ctx context.Context, request *workflow.Records) (*workflow.Response, error) {
	records := request.GetEntries();
	err := database.ChangeStatusTo(status.COMPLETED, records);

	message := fmt.Sprintf("Successfully changed status of %v to completed", records);
	if (err != nil) {
		message = fmt.Sprintf("Failed to change status of %v to completed", records);
	}
	return &workflow.Response{Message: message}, err
}