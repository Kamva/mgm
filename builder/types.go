// Package builder help us to write aggregate,filter,update maps simpler.
package builder

import "go.mongodb.org/mongo-driver/bson"

// SMap is simple map that can be substitute  of `bson.M` to
// having simpler map structure on query,aggregate,...
type SMap struct {
	Operators []Operator
}

// ToMap function convert our Simple map to bson.M to using in filters,stages,...
func (s *SMap) ToMap() bson.M {
	m := bson.M{}

	for _, o := range s.Operators {
		m[o.GetKey()] = o.GetVal()
	}

	return m
}

// S get operators as param and return bson.M to using result as filter,stage,...
func S(operators ...Operator) bson.M {
	s := &SMap{Operators: operators}

	return s.ToMap()
}
