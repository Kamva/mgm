package crud

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}
