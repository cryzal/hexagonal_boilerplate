package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	logKey "hexagonal_boilerplate/shared/utils/logger/libs/const/log"
	"hexagonal_boilerplate/shared/utils/logger/log"
	"time"
)

type (
	RepositoryDbtransaction struct {
		*mongo.Client
	}
)

func NewDbtransactionRepo(db *mongo.Client) *RepositoryDbtransaction {
	return &RepositoryDbtransaction{db}
}

var txKey *mongo.Client

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

	session, err := repo.StartSession()
	if err != nil {
		return err
	}
	sessionCtx := mongo.NewSessionContext(ctx, session)

	err = session.StartTransaction()
	if err != nil {
		panic(err)
	}

	// run callback
	err = tFunc(sessionCtx)
	if err != nil {
		errRollback := mongo.SessionFromContext(sessionCtx).AbortTransaction(sessionCtx)

		if errRollback != nil {
			return errRollback
		}

		mongo.SessionFromContext(sessionCtx).EndSession(sessionCtx)
		return err
	}

	err = mongo.SessionFromContext(sessionCtx).CommitTransaction(sessionCtx)
	if err != nil {
		return err
	}

	mongo.SessionFromContext(sessionCtx).EndSession(sessionCtx)
	return nil
}
