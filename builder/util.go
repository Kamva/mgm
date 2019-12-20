package builder

import (
	"github.com/Kamva/mgm/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func appendIfIsNotEmpty(m bson.M, key string, val interface{}) {
	if strVal, ok := val.(string); ok {
		if strVal = strings.Trim(strVal, " "); len(strVal) != 0 {
			m[key] = strVal
		}

		return
	}

	if !util.InterfaceIsNil(val) {
		m[key] = val
	}
}
