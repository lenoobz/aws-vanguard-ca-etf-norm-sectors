package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardOverviewModel represents Vanguard fund overview model
type VanguardOverviewModel struct {
	ID               *primitive.ObjectID      `bson:"_id,omitempty"`
	IsActive         bool                     `bson:"isActive,omitempty"`
	CreatedAt        int64                    `bson:"createdAt,omitempty"`
	ModifiedAt       int64                    `bson:"modifiedAt,omitempty"`
	Schema           string                   `bson:"schema,omitempty"`
	PortID           string                   `bson:"portId,omitempty"`
	AssetClass       string                   `bson:"assetClass,omitempty"`
	Strategy         string                   `bson:"strategy,omitempty"`
	DividendSchedule string                   `bson:"dividendSchedule,omitempty"`
	Name             string                   `bson:"name,omitempty"`
	Currency         string                   `bson:"currency,omitempty"`
	Isin             string                   `bson:"isin,omitempty"`
	Sedol            string                   `bson:"sedol,omitempty"`
	Ticker           string                   `bson:"ticker,omitempty"`
	TotalAssets      float64                  `bson:"totalAssets,omitempty"`
	Yield12Month     float64                  `bson:"yield12Month,omitempty"`
	Price            float64                  `bson:"price,omitempty"`
	ManagementFee    float64                  `bson:"managementFee,omitempty"`
	MerFee           float64                  `bson:"merFee,omitempty"`
	Sectors          []*SectorBreakdownModel  `bson:"sectors,omitempty"`
	Countries        []*CountryBreakdownModel `bson:"countries,omitempty"`
}

// SectorBreakdownModel is the representation of sector the fund invested
type SectorBreakdownModel struct {
	FundPercent float64 `bson:"fundPercent,omitempty"`
	SectorName  string  `bson:"sectorName,omitempty"`
}

// CountryBreakdownModel is the representation of country the fund exposed
type CountryBreakdownModel struct {
	CountryName     string  `bson:"countryName,omitempty"`
	FundTnaPercent  float64 `bson:"fundTnaPercent,omitempty"`
	FundMktPercent  float64 `bson:"fundMktPercent,omitempty"`
	HoldingStatCode string  `bson:"holdingStatCode,omitempty"`
}
