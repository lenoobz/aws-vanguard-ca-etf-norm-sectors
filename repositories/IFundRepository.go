package repositories

import "github.com/hthl85/aws-vanguard-ca-etf-sectors/repositories/mongodb/models"

// IFundRepository interface
type IFundRepository interface {
	GetAllFundsOverview() ([]*models.FundOverviewModel, error)
	UpdateAllFundsOverview([]*models.FundOverviewModel) error
}
