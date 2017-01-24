package database

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"fmt"
	"log"
	"Retail/workflow/status"
	workFlow "github.com/RetailMarket/workFlowClient"
)

const (
	DB_DRIVER = "postgres"
	DB_CONNECTION = "user=postgres dbname=postgres password=postgres sslmode=disable"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open(DB_DRIVER, DB_CONNECTION)
	if (err != nil) {
		log.Fatal(err.Error())
	}
	db.Ping()
}

func CloseDb() {
	db.Close()
}

func SavePriceInUpdateApprovalTable(priceList []*workFlow.Product) error {
	fmt.Println("Inserting record into updatePriceApproval table...")
	insertQuery := "insert into workflow.priceUpdateApproval (product_id, product_name,cost,status) values($1,$2,$3,$4)"

	for i := 0; i < len(priceList); i++ {
		priceObj := priceList[i];
		_, err := db.Exec(
			insertQuery,
			priceObj.ProductId,
			priceObj.ProductName,
			priceObj.Cost,
			priceObj.Status)
		if (err != nil) {
			return err;
			//log.Fatalf("Unable to save update entry in db for approval \nError: %v", err.Error());
		}
	}
	return nil;
}

func FetchPriceUpdateRequests() {
	fmt.Println("Inserting record into updatePriceApproval table...")
	selectQuery := "select * from price.updateApproval where status = $1"

	_, err := db.Query(
		selectQuery, status.PENDING)

	if (err != nil) {
		log.Fatalf("Unable to fetch update requests from db \nError: %v", err.Error());
	}
}