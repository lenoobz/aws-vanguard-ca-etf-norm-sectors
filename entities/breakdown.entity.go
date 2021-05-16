package entities

// FundBreakdown represents Vanguard's fund sector
type FundBreakdown struct {
	Ticker     string             `json:"ticker,omitempty"`
	AssetClass string             `json:"assetClass,omitempty"`
	Sectors    []*SectorBreakdown `json:"sectors,omitempty"`
}

// SectorBreakdown is the representation of sector the fund invested
type SectorBreakdown struct {
	SectorCode  string  `json:"sectorCode,omitempty"`
	SectorName  string  `json:"sectorName,omitempty"`
	FundPercent float64 `json:"fundPercent,omitempty"`
}
