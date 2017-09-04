package main

import (
	"fmt"
	"os"
)

func main() {
	db := getDatabase(os.Getenv("DATABASE_URL"))
	if os.Getenv("MIGRATE_DB") == "true" {
		migrate(db)
	}
	defer db.Close()

	router := initializeRoutes(db)
	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%v", port))
}
