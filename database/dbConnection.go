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
	SCHEMA = "workflow"
	UPDATE_APPROVAL_TABLE = "priceUpdateApproval"
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

func SavePriceInUpdateApprovalTable(priceList []*workFlow.Entry) error {
	fmt.Println("Inserting Record Into UpdatePriceApproval Table...")

	for i := 0; i < len(priceList); i++ {
		priceObj := priceList[i];
		insertQuery := fmt.Sprintf("insert into %s.%s (product_id, version, status) values(%d,'%s','%s')",
			SCHEMA,
			UPDATE_APPROVAL_TABLE,
			priceObj.ProductId,
			priceObj.Version,
			status.PENDING)

		_, err := db.Exec(
			insertQuery)
		if (err != nil) {
			log.Fatalf("Unable to save update entry in db for approval \nError: %v", err.Error());
			return err;
		}
	}
	return nil;
}

func GetAllPendingRecords() (*sql.Rows, error) {
	fmt.Println("Fetching Pending Record From UpdatePriceApproval Table...")
	selectQuery := fmt.Sprintf("select product_id, version from %s.%s where status = '%s'", SCHEMA, UPDATE_APPROVAL_TABLE, status.PENDING)
	return db.Query(selectQuery)
}

func ChangeStatusTo(status string, records []*workFlow.Entry) error {
	for i := 0; i < len(records); i++ {
		updateQuery := fmt.Sprintf("update %s.%s set status = '%s' where product_id = %d and version = '%s'", SCHEMA, UPDATE_APPROVAL_TABLE, status, int(records[i].ProductId), records[i].Version)
		_, err := db.Exec(updateQuery)
		if (err != nil) {
			return err;
		}
	}
	return nil;
}