package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VanguardOverviewModel struct
type VanguardOverviewModel struct {
	ID               *primitive.ObjectID     `bson:"_id,omitempty"`
	IsActive         bool                    `bson:"isActive,omitempty"`
	CreatedAt        int64                   `bson:"createdAt,omitempty"`
	ModifiedAt       int64                   `bson:"modifiedAt,omitempty"`
	Schema           string                  `bson:"schema,omitempty"`
	PortID           string                  `bson:"portId,omitempty"`
	AssetClass       string                  `bson:"assetClass,omitempty"`
	Strategy         string                  `bson:"strategy,omitempty"`
	DividendSchedule string                  `bson:"dividendSchedule,omitempty"`
	Name             string                  `bson:"name,omitempty"`
	ShortName        string                  `bson:"shortName,omitempty"`
	Currency         string                  `bson:"currency,omitempty"`
	Isin             string                  `bson:"isin,omitempty"`
	Sedol            string                  `bson:"sedol,omitempty"`
	Ticker           string                  `bson:"ticker,omitempty"`
	TotalAssets      float64                 `bson:"totalAssets,omitempty"`
	Yield12Month     float64                 `bson:"yield12Month,omitempty"`
	Price            float64                 `bson:"price,omitempty"`
	ManagementFee    float64                 `bson:"managementFee,omitempty"`
	MerFee           float64                 `bson:"merFee,omitempty"`
	DistYield        float64                 `bson:"distYield,omitempty"`
	AllocationStock  float64                 `bson:"allocationStock,omitempty"`
	AllocationBond   float64                 `bson:"allocationBond,omitempty"`
	AllocationCash   float64                 `bson:"allocationCash,omitempty"`
	Sectors          []*SectorBreakdownModel `bson:"sectors,omitempty"`
}

// SectorBreakdownModel struct
type SectorBreakdownModel struct {
	SectorCode  string  `bson:"sectorCode,omitempty"`
	SectorName  string  `bson:"sectorName,omitempty"`
	FundPercent float64 `bson:"fundPercent,omitempty"`
}
