package mgm_test

import (
	"github.com/Kamva/mgm/v2"
	"github.com/Kamva/mgm/v2/internal/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

func TestSetupDefaultConnection(t *testing.T) {
	err := mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))

	require.Nil(t, err)
}

func TestSetupWrongConnection(t *testing.T) {
	err := mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("wrong://wrong:wrong@localhost:27017"))

	require.NotNil(t, err)
}

func TestPanicOnGetCtx(t *testing.T) {
	mgm.ResetDefaultConfig()

	defer func() {
		require.NotNil(t, recover(), "Getting context before set default config must panic")
	}()

	_ = mgm.Ctx()
}

func TestGetCtx(t *testing.T) {
	defer func() {
		require.Nil(t, recover(), "Getting context after set default config must return a context.")
	}()

	// Setup connection
	setupDefConnection()

	ctx := mgm.Ctx()

	_, ok := ctx.Deadline()
	require.True(t, ok, "context should having deadline.")
}

func TestGetCollection(t *testing.T) {
	// Setup connection
	setupDefConnection()

	col := mgm.CollectionByName("test_collection")

	require.Equal(t, col.Name(), "test_collection")
}

func TestGetNewClient(t *testing.T) {
	// Setup default config:
	setupDefConnection()

	client, err := mgm.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	util.AssertErrIsNil(t, err)

	// Check client connection:
	err = client.Ping(mgm.Ctx(), readpref.Primary())
	util.AssertErrIsNil(t, err)

	// Get New Collection:
	coll := mgm.NewCollection(client.Database("test_db"), "test_col")

	require.Equal(t, coll.Name(), "test_col", "expected collection with %v name, got %v", "test_col", coll.Name())
}

func TestGetDefaultConfigBeforeSettingItUp(t *testing.T) {
	mgm.ResetDefaultConfig()

	_, _, _, err := mgm.DefaultConfigs()
	require.NotNil(t, err, "Expected get error when getting config before setting up it.")
}

func TestGetDefaultConfigAfterSettingItUp(t *testing.T) {
	setupDefConnection()

	conf, client, db, err := mgm.DefaultConfigs()

	if util.AnyNil(conf, client, db) {
		t.Errorf("expired get config,client,db after setting up default config, got %v,%v,%v", conf, client, db)
	}

	util.AssertErrIsNil(t, err)
}
