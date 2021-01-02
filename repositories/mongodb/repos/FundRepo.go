package repos

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hthl85/aws-vanguard-ca-etf-sectors/consts"
	"github.com/hthl85/aws-vanguard-ca-etf-sectors/repositories/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// FundRepo struct
type FundRepo struct {
	DB *mongo.Database
}

// NewFundRepo creates new fund mongo repo
func NewFundRepo(db *mongo.Database) (*FundRepo, error) {
	fmt.Println("Create new Fund Repo")

	if db != nil {
		return &FundRepo{
			DB: db,
		}, nil
	}

	// set context with timeout from the config
	timeout := time.Duration(consts.TimeoutMS) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// set mongo client options
	clientOptions := options.Client()

	// set min pool size
	if consts.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(consts.MinPoolSize)
	}

	// set max pool size
	if consts.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(consts.MaxPoolSize)
	}

	// set max idle time ms
	if consts.MaxIdleTimeMS > 0 {
		clientOptions.SetMaxConnIdleTime(time.Duration(consts.MaxIdleTimeMS) * time.Millisecond)
	}

	// construct a connection string from mongo config object
	cxnString := fmt.Sprintf("mongodb+srv://%s:%s@%s", consts.Username, consts.Password, consts.Host)

	// create mongo client by making new connection
	client, err := mongo.Connect(ctx, clientOptions.ApplyURI(cxnString))
	if err != nil {
		return nil, err
	}

	// test our connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &FundRepo{
		DB: client.Database(consts.Dbname),
	}, nil
}

///////////////////////////////////////////////////////////////////////////////
// Implement interface
///////////////////////////////////////////////////////////////////////////////

// GetAllFundsOverview find fund country exposure
func (repo *FundRepo) GetAllFundsOverview() ([]*models.FundOverviewModel, error) {
	fmt.Println("Get All Funds Overview")

	// create new context for the query
	ctx, cancel := createContext()
	defer cancel()

	// what collection we are going to use
	col := repo.DB.Collection(consts.CollectionVanguardFundOverview)

	// filter
	filter := bson.D{}

	// find options
	findOptions := options.Find()

	cur, err := col.Find(ctx, filter, findOptions)

	// only run defer function when find success
	if cur != nil {
		defer func() {
			if deferErr := cur.Close(ctx); deferErr != nil {
				err = deferErr
			}
		}()
	}

	// find was not succeed
	if err != nil {
		return nil, err
	}

	var funds []*models.FundOverviewModel

	// iterate over the cursor to decode document one at a time
	for cur.Next(ctx) {
		// decode cursor to activity model
		var fundOverviewModel models.FundOverviewModel
		if err = cur.Decode(&fundOverviewModel); err != nil {
			fmt.Println("error decode fund overview", err.Error())
			return nil, err
		}

		if fundOverviewModel.SectorWeighting != nil {
			funds = append(funds, &fundOverviewModel)
		}
	}

	if err := cur.Err(); err != nil {
		fmt.Println("error iterate over the cursor")
		return nil, err
	}

	return funds, nil
}

// UpdateAllFundsOverview find fund country exposure
func (repo *FundRepo) UpdateAllFundsOverview(funds []*models.FundOverviewModel) error {
	fmt.Println("Update All Funds Overview")

	// create new context for the query
	ctx, cancel := createContext()
	defer cancel()

	// what collection we are going to use
	col := repo.DB.Collection(consts.CollectionFundSectors)

	for _, fund := range funds {
		fund.ModifiedAt = time.Now().UTC().Unix()

		filter := bson.D{{
			Key:   "ticker",
			Value: fund.Ticker,
		}}

		update := bson.D{{
			Key:   "$set",
			Value: fund,
		}}

		opts := options.Update().SetUpsert(true)

		if _, err := col.UpdateOne(ctx, filter, update, opts); err != nil {
			fmt.Println("error update fund with ticker", fund.Ticker)
			return err
		}

	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////
// Private helper functions
///////////////////////////////////////////////////////////////////////////////

// createContext create a new context with timeout
func createContext() (context.Context, context.CancelFunc) {
	// set context with timeout from the config
	timeout := time.Duration(consts.TimeoutMS) * time.Millisecond
	return context.WithTimeout(context.Background(), timeout*time.Millisecond)
}

func getCountryCode(name string, countryConsts []models.Country) string {
	for _, v := range countryConsts {
		if strings.ToUpper(v.Name) == strings.ToUpper(name) {
			return v.Code
		}
	}
	return ""
}
