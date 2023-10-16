package aggregate

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "mgmdb", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}
