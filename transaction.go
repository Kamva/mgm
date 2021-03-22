package mgm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// TransactionFunc is a handler to manage a transaction.
type TransactionFunc func(session mongo.Session, sc mongo.SessionContext) error

// Transaction creates a transaction with the default client.
func Transaction(f TransactionFunc) error {
	return TransactionWithClient(ctx(), client, f)
}

// TransactionWithCtx creates a transaction with the given context and the default client.
func TransactionWithCtx(ctx context.Context, f TransactionFunc) error {
	return TransactionWithClient(ctx, client, f)
}

// TransactionWithClient creates a transaction with the given client.
func TransactionWithClient(ctx context.Context, client *mongo.Client, f TransactionFunc) error {
	session, err := client.StartSession() //start session need to get options.
	if err != nil {
		return err
	}

	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil { // startTransaction need to get options.
		return err
	}

	wrapperFn := func(sc mongo.SessionContext) error {
		return f(session, sc)
	}

	return mongo.WithSession(ctx, session, wrapperFn)
}
