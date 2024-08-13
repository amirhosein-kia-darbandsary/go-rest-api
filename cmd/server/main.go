package main

import (
	"context"
	"fmt"

	"github.com/amirhosein-kia-darbandsary/project/internal/comment"
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
	cmtService := comment.NewService(db)

	cmtService.DeleteComment(context.Background(), "43edc8d0-6d3c-457d-81ad-1e2473aebc92")

}

func main() {
	Run()
}
