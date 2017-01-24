package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"Retail/workflow/database"
)

func GetAllUpdateApprovalRequest(c *gin.Context) {
	database.FetchPriceUpdateRequests();

	updatedPrice := model.Price{}
	updatedPrice.Product_id = id
	updatedPrice.Product_name = name
	updatedPrice.Cost = price
	updatedPrice.Status = status.PENDING

	database.SavePriceInUpdateTable(db, &updatedPrice);
}
