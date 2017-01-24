package routes

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	"Retail/workflow/handler"
)

func HandleRequest(db *sql.DB) {
	router := gin.Default()

	router.POST("/requests/update", handler.GetAllUpdateApprovalRequest)


	router.Handle("/", http.FileServer(http.Dir("./public")));
	//handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", router)

	router.Run(":4000")
}

