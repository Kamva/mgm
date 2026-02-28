package mgm

import (
	"reflect"

	"github.com/jinzhu/inflection"
	"github.com/kamva/mgm/v3/internal/util"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Coll returns the collection associated with a model.
func Coll(m Model, opts ...options.Lister[options.CollectionOptions]) *Collection {

	if collGetter, ok := m.(CollectionGetter); ok {
		return collGetter.Collection()
	}

	return CollectionByName(CollName(m), opts...)
}

// CollName returns a model's collection name. The `CollectionNameGetter` will be used
// if the model implements this interface. Otherwise, the collection name is inferred
// based on the model's type using reflection.
func CollName(m Model) string {

	if collNameGetter, ok := m.(CollectionNameGetter); ok {
		return collNameGetter.CollectionName()
	}

	name := reflect.TypeOf(m).Elem().Name()

	return inflection.Plural(util.ToSnakeCase(name))
}

// UpsertTrueOption returns a new UpdateOne options builder with the upsert property set to true.
func UpsertTrueOption() *options.UpdateOneOptionsBuilder {
	return options.UpdateOne().SetUpsert(true)
}
