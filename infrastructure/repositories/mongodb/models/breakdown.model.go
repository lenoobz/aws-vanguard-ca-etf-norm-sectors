package models

import (
	"github.com/hthl85/aws-vanguard-ca-etf-norm-sectors/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundBreakdownModel is the representation of individual Vanguard fund overview model
type FundBreakdownModel struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	IsActive   bool                `bson:"isActive,omitempty"`
	CreatedAt  int64               `bson:"createdAt,omitempty"`
	ModifiedAt int64               `bson:"modifiedAt,omitempty"`
	Schema     string              `bson:"schema,omitempty"`
	Ticker     string              `bson:"ticker,omitempty"`
	AssetClass string              `bson:"assetClass,omitempty"`
	Sectors    []*BreakdownModel   `bson:"sectors,omitempty"`
}

// BreakdownModel is the representation of country the fund exposed
type BreakdownModel struct {
	SectorCode  string  `bson:"sectorCode,omitempty"`
	SectorName  string  `bson:"sectorName,omitempty"`
	FundPercent float64 `bson:"fundPercent,omitempty"`
}

// NewFundBreakdownModel create new fund exposure model
func NewFundBreakdownModel(e *entities.FundBreakdown) *FundBreakdownModel {
	var m []*BreakdownModel

	for _, v := range e.Sectors {
		m = append(m, &BreakdownModel{
			FundPercent: v.FundPercent,
			SectorName:  v.SectorName,
			SectorCode:  v.SectorCode,
		})
	}

	return &FundBreakdownModel{
		Ticker:     e.Ticker,
		AssetClass: e.AssetClass,
		Sectors:    m,
	}
}
