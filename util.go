package mgm

import (
	"github.com/Kamva/mgm/internal/util"
	"github.com/jinzhu/inflection"
	"reflect"
)

// Coll return model's collection.
func Coll(m Model) *Collection {

	if collGetter, ok := m.(CollectionGetter); ok {
		return collGetter.Collection()
	}

	return CollectionByName(CollName(m))
}

// CollName check if you provided collection name in your
// model, return it's name, otherwise guess model
// collection's name.
func CollName(m Model) string {

	if collNameGetter, ok := m.(CollectionNameGetter); ok {
		return collNameGetter.CollectionName()
	}

	name := reflect.TypeOf(m).Elem().Name()

	return inflection.Plural(util.ToSnakeCase(name))
}
