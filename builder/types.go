// builder package help us to write aggregate,filter,update maps simpler.
package builder

import "go.mongodb.org/mongo-driver/bson"

// SMap is simple map that can be substitute  of `bson.M` to
// having simpler map structure on query,aggregate,...
type SMap struct {
	Operators []Operator
}

func (s *SMap) ToMap() bson.M {
	m := bson.M{}

	for _, o := range s.Operators {
		m[o.GetKey()] = o.GetVal()
	}

	return m
}

func S(operators ...Operator) bson.M {
	s := &SMap{Operators: operators}

	return s.ToMap()
}
