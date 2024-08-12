package main

import (
	"context"
	"fmt"

	"github.com/amirhosein-kia-darbandsary/project/internal/db"
)

func Run() {
	db, err := db.NewDatabaseConnection()
	if err != nil {
		fmt.Println("Db has Not connection Detail :", err)
	}
	pinged := db.PingDatabase(context.Background())
	if pinged != nil {
		fmt.Println("ERROR IN PINGING THE DB....")
	}
	fmt.Println("STarting migrating .......")
	db.Migrate()
	fmt.Println("Run the server......")
}

func main() {
	Run()
}
