package main

import (
	"Retail/workflow/database"
	"Retail/workflow/server"
)

func main() {
	database.Init();
	defer database.CloseDb();

	server.CreateServerConnection();
}