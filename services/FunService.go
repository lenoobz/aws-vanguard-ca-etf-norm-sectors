package services

import (
	"fmt"

	"github.com/hthl85/aws-vanguard-ca-etf-sectors/repositories"
)

// FundService struct
type FundService struct {
	fundRepo repositories.IFundRepository
}

// NewFundService create as new service
func NewFundService(fundRepo repositories.IFundRepository) *FundService {
	fmt.Println("Create new Fund Service")

	return &FundService{
		fundRepo: fundRepo,
	}
}

// PopulateFundSectors find fund country exposure
func (svc *FundService) PopulateFundSectors() error {
	fmt.Println("Populate Fund Countries")

	funds, err := svc.fundRepo.GetAllFundsOverview()
	if err != nil {
		return err
	}

	return svc.fundRepo.UpdateAllFundsOverview(funds)
}
