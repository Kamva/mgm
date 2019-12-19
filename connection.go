package mgm

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mgm/internal"
	"time"
)

var config *Config
var client *mongo.Client
var db *mongo.Database

type Config struct {
	// Set to 10 second (10*time.Second) for example.
	CtxTimeout time.Duration
}

func NewCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	return ctx
}

func Ctx() context.Context {
	return ctx()
}

func ctx() context.Context {
	return NewCtx(config.CtxTimeout)
}

func NewClient(opts ...*options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	if err = client.Connect(Ctx()); err != nil {
		return nil, err
	}

	return client, nil
}

// NewCollection return new collection with passed database
func NewCollection(db *mongo.Database, name string, opts ...*options.CollectionOptions) *Collection {
	coll := db.Collection(name, opts...)

	return &Collection{Collection: coll}
}

// ResetDefaultConfig reset all of the default config
func ResetDefaultConfig() {
	config = nil
	client = nil
	db = nil
}

// SetDefaultConfig initial default client and Database .
func SetDefaultConfig(conf *Config, dbName string, opts ...*options.ClientOptions) (err error) {

	// Get predefined config as default config if user
	// do not provide it.
	if conf == nil {
		conf = defaultConf()
	}

	config = conf

	if client, err = NewClient(opts...); err != nil {
		return err
	}

	db = client.Database(dbName)

	return nil
}

// GetCollection return new collection from default config
func GetCollection(name string, opts ...*options.CollectionOptions) *Collection {
	return NewCollection(db, name)
}

func DefaultConfigs() (*Config, *mongo.Client, *mongo.Database, error) {
	if internal.AnyNil(config, client, db) {
		return nil, nil, nil, errors.New("please setup default config before acquiring it")
	}

	return config, client, db, nil
}

// defaultConf is default config ,If you do not pass config
// to `SetDefaultConfig` method, we using this config.
func defaultConf() *Config {
	return &Config{CtxTimeout: 10 * time.Second}
}
