// Package builder help us to write aggregates, filters, update maps simpler.
package builder

import "go.mongodb.org/mongo-driver/bson"

// SMap is simple map that can be substitute of `bson.M` to
// have a simpler map structure for queries, aggregations, etc.
type SMap struct {
	Operators []Operator
}

// ToMap function converts our SMap to bson.M for use in filters, stages, etc.
func (s *SMap) ToMap() bson.M {
	m := bson.M{}

	for _, o := range s.Operators {
		m[o.GetKey()] = o.GetVal()
	}

	return m
}

// S receives operators as parameters and returns a bson.M that can be used in filters, stages, etc.
func S(operators ...Operator) bson.M {
	s := &SMap{Operators: operators}

	return s.ToMap()
}
