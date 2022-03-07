package postgres

import (
	"context"
	core "github.com/Qalifah/grey-challenge/wallet"
	"github.com/jackc/pgx/v4"
)

type balanceRepository struct {
	client *pgx.Conn
}

func NewBalanceRepository(client *pgx.Conn) *balanceRepository {
	return &balanceRepository{
		client: client,
	}
}

func (r *balanceRepository) Get(ctx context.Context, userID int) (*core.Balance, error) {
	row := r.client.QueryRow(ctx, "SELECT * FROM balances WHERE user_id = $1", userID)
	balance := &core.Balance{}
	err := row.Scan(&balance.ID, &balance.UserID, &balance.Amount, &balance.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (r *balanceRepository) Update(ctx context.Context, userID, newAmount int) (*core.Balance, error) {
	cmd := r.client.QueryRow(ctx, "UPDATE balances SET updated_at = CURRENT_TIMESTAMP, amount = $1 WHERE user_id = $2 RETURNING id, user_id, amount, updated_at", newAmount, userID)
	bal := &core.Balance{}
	err := cmd.Scan(&bal.ID, &bal.UserID, &bal.Amount, &bal.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return bal, nil
}
