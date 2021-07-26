package breakdown

import (
	"context"

	"github.com/lenoobz/aws-vanguard-ca-etf-norm-sectors/entities"
)

///////////////////////////////////////////////////////////
// Sector Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface {
	FindSectorsBreakdown(context.Context) ([]*entities.FundBreakdown, error)
}

// Writer interface
type Writer interface {
	UpdateSectorsBreakdown(context.Context, []*entities.FundBreakdown) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
