package mgm_test

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mgm"
	"mgm/internal"
	"testing"
)

func setupDefConnection() {
	internal.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("mongodb://root:12345@localhost:27017")),
	)
}

func resetCollection() {
	_, err := mgm.ModelCollection(&Doc{}).DeleteMany(mgm.Ctx(), bson.M{})
	_, err2 := mgm.ModelCollection(&Person{}).DeleteMany(mgm.Ctx(), bson.M{})

	internal.PanicErr(err)
	internal.PanicErr(err2)
}

func seed() {
	docs := []interface{}{
		NewDoc("Ali", 24),
		NewDoc("Mehran", 25),
		NewDoc("Reza", 26),
		NewDoc("Omid", 27),
	}
	_, err := mgm.ModelCollection(&Doc{}).InsertMany(mgm.Ctx(), docs)

	internal.PanicErr(err)
}

func findDoc(t *testing.T) *Doc {
	found := &Doc{}
	internal.AssertErrIsNil(t, mgm.ModelCollection(found).FindOne(mgm.Ctx(), bson.M{}).Decode(found))

	return found
}

type Doc struct {
	mgm.DefaultModel `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (d *Doc) Collection() *mgm.Collection {
	return mgm.GetCollection("docs")
}

func NewDoc(name string, age int) *Doc {
	return &Doc{Name: name, Age: age}
}
