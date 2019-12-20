package builder

import (
	"github.com/Kamva/mgm/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

// appendIfHasVal append key and val to map if value is not empty.
func appendIfHasVal(m bson.M, key string, val interface{}) {
	if !util.IsNil(val) {
		m[key] = val
	}
}
