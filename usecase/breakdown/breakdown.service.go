package breakdown

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
)

// Service exposure
type Service struct {
	repo Repo
	log  logger.ContextLog
}

// NewService create new service
func NewService(r Repo, l logger.ContextLog) *Service {
	return &Service{
		repo: r,
		log:  l,
	}
}

// PopulateFundSectors populate fund sectors data
func (s *Service) PopulateFundSectors(ctx context.Context) error {
	s.log.Info(ctx, "populate fund sector weighted")

	sectors, err := s.repo.FindSectorsBreakdown(ctx)
	if err != nil {
		s.log.Error(ctx, "find all sectors failed", "error", err)
		return err
	}

	if err := s.repo.UpdateSectorsBreakdown(ctx, sectors); err != nil {
		s.log.Error(ctx, "update all sectors failed", "error", err)
		return err
	}

	return nil
}
