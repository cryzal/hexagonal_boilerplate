package mysql

import (
	"context"
	logKey "hexagonal_boilerplate/shared/utils/logger/libs/const/log"
	"hexagonal_boilerplate/shared/utils/logger/log"
	"time"

	"gorm.io/gorm"
)

type (
	RepositoryDbtransaction struct {
		*gorm.DB
	}
)

func NewDbtransactionRepo(db *gorm.DB) *RepositoryDbtransaction {
	return &RepositoryDbtransaction{db}
}

var txKey *gorm.DB

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {

		return tx
	}
	return nil
}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, tx)
}
func (repo *RepositoryDbtransaction) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	startTime := time.Now()
	// register logs
	log := log.With().
		Str(logKey.SERVICE_NAME, logKey.TRANSACTION_REPOSITORY).
		Str(logKey.SERVICE_METHOD, "WithinTransaction").
		Time(logKey.START_TIME, startTime).
		Logger()

	log.Info().
		Msg("")

	tx := repo.DB.Begin()
	log.Info().
		Msg("")

	// run callback
	err := tFunc(injectTx(ctx, tx))
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Printf("rollback transaction: %v", errRollback)
		}
		return err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		log.Printf("commit transaction: %v", errCommit)
	}
	return nil
}
