package mgm_test

import (
	"github.com/Kamva/mgm/v2"
	"github.com/Kamva/mgm/v2/internal/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

// Note: to run Transaction tests:
// - MongoDB daemon must run as replica set not standalone
//	 daemon, to convert it [see this](https://docs.mongodb.com/manual/tutorial/convert-standalone-to-replica-set/)
func TestTransactionCommit(t *testing.T) {
	setupDefConnection()
	resetCollection()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).SaveWithCtx(sc, d)

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
	//seed()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).SaveWithCtx(sc, d)

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
	//seed()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.TransactionWithCtx(mgm.Ctx(), func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).SaveWithCtx(sc, d)

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
