package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hthl85/aws-vanguard-ca-etf-sectors/repositories/mongodb/repos"
	"github.com/hthl85/aws-vanguard-ca-etf-sectors/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func main() {
	repo, err := repos.NewFundRepo(db)
	if err != nil {
		fmt.Println("error occurred when connect to database", err)
	}

	// we won't close database connection
	db = repo.DB

	// init service
	svc := services.NewFundService(repo)

	// if err = svc.PopulateFundSectors(); err != nil {
	// 	fmt.Println("error populate fund sectors")
	// }
	lambda.Start(svc.PopulateFundSectors)
}
