package core

import (
	"context"
	"time"
)

type Balance struct {
	ID        int
	UserID    int
	Amount    int
	UpdatedAt time.Time
}

type BalanceRepository interface {
	Get(ctx context.Context, userID int) (*Balance, error)
	Update(ctx context.Context, userID, newAmount int) (*Balance, error)
}
