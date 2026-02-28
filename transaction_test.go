package mgm_test

import (
	"context"
	"errors"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/internal/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"testing"
)

// Note: to run Transaction tests, the MongoDB daemon must run as replica set, not as a standalone daemon.
// To convert it [see this](https://docs.mongodb.com/manual/tutorial/convert-standalone-to-replica-set/)
func TestTransactionCommit(t *testing.T) {
	setupDefConnection()
	resetCollection()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session *mongo.Session, sc context.Context) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.CommitTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(1), count)
}

func TestTransactionAbort(t *testing.T) {
	setupDefConnection()
	resetCollection()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session *mongo.Session, sc context.Context) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.AbortTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}

func TestTransactionWithCtx(t *testing.T) {
	setupDefConnection()
	resetCollection()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.TransactionWithCtx(mgm.Ctx(), func(session *mongo.Session, sc context.Context) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.AbortTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}

func TestTransactionErrorPropagation(t *testing.T) {
	setupDefConnection()
	resetCollection()

	expectedErr := errors.New("transaction error")

	err := mgm.Transaction(func(session *mongo.Session, sc context.Context) error {
		return expectedErr
	})

	require.Error(t, err)
	require.ErrorIs(t, err, expectedErr)

	// Collection should be empty since nothing was inserted
	count, err := mgm.Coll(&Doc{}).CountDocuments(mgm.Ctx(), bson.M{})
	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}

func TestTransactionCommitMultipleOps(t *testing.T) {
	setupDefConnection()
	resetCollection()

	err := mgm.Transaction(func(session *mongo.Session, sc context.Context) error {
		d1 := &Doc{Name: "doc1", Age: 10}
		d2 := &Doc{Name: "doc2", Age: 20}

		if err := mgm.Coll(d1).CreateWithCtx(sc, d1); err != nil {
			return err
		}
		if err := mgm.Coll(d2).CreateWithCtx(sc, d2); err != nil {
			return err
		}

		return session.CommitTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(&Doc{}).CountDocuments(mgm.Ctx(), bson.M{})
	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(2), count)
}

func TestTransactionWithClientExplicit(t *testing.T) {
	setupDefConnection()
	resetCollection()

	_, cl, _, err := mgm.DefaultConfigs()
	util.AssertErrIsNil(t, err)

	d := &Doc{Name: "check", Age: 10}

	err = mgm.TransactionWithClient(mgm.Ctx(), cl, func(session *mongo.Session, sc context.Context) error {
		if err := mgm.Coll(d).CreateWithCtx(sc, d); err != nil {
			return err
		}
		return session.CommitTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})
	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(1), count)
}
