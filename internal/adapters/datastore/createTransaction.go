package datastore

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

func (ds *Datastore) CreateTransaction(ctx context.Context, transaction entities.Transaction) error {
	_, err := ds.db.ExecContext(
		ctx,
		"INSERT INTO transactions (id,account_id,operation_id,amount,event_time) VALUES ($1,$2,$3,$4,$5)",
		transaction.ID,
		transaction.AccountID,
		transaction.OperationType,
		transaction.Amount,
		transaction.EventDate,
	)
	if err != nil {
		ds.logger.Error("failed to create transaction in database", "error", err.Error())
	}
	return err
}
