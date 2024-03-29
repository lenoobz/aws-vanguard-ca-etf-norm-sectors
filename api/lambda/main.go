package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-vanguard-ca-etf-norm-sectors/config"
	"github.com/lenoobz/aws-vanguard-ca-etf-norm-sectors/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-vanguard-ca-etf-norm-sectors/usecase/breakdown"
)

func main() {
	appConf := config.AppConf

	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	// create new repository
	repo, err := repos.NewBreakdownMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create fund mongo repo failed")
	}
	defer repo.Close()

	// create new service
	svc := breakdown.NewService(repo, zap)

	lambda.Start(svc.PopulateFundSectors)
}
