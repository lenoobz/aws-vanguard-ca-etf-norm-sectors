package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Country struct
type Country struct {
	Name string `json:"name"`
	Code string `json:"alpha3Code"`
}

// FundOverviewModel is the representation of individual Vanguard fund overview model
type FundOverviewModel struct {
	ID              *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Schema          int                 `json:"schema,omitempty" bson:"schema,omitempty"`
	IsActive        bool                `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAt       int64               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt      int64               `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
	Ticker          string              `json:"ticker,omitempty" bson:"ticker,omitempty"`
	AssetClass      string              `json:"assetClass,omitempty" bson:"assetClass,omitempty"`
	SectorWeighting []*SectorWeighting  `json:"sectorWeighting,omitempty"`
}

// SectorWeighting is the representation of sector the fund invested
type SectorWeighting struct {
	FundPercent float64 `json:"fundPercent,omitempty"`
	LongName    string  `json:"longName,omitempty"`
	SectorType  string  `json:"sectorType,omitempty"`
}
