package mgm_test

import (
	"github.com/Kamva/mgm"
	"github.com/Kamva/mgm/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func setupDefConnection() {
	util.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("mongodb://root:12345@localhost:27017")),
	)
}

func resetCollection() {
	_, err := mgm.Coll(&Doc{}).DeleteMany(mgm.Ctx(), bson.M{})
	_, err2 := mgm.Coll(&Person{}).DeleteMany(mgm.Ctx(), bson.M{})

	util.PanicErr(err)
	util.PanicErr(err2)
}

func seed() {
	docs := []interface{}{
		NewDoc("Ali", 24),
		NewDoc("Mehran", 25),
		NewDoc("Reza", 26),
		NewDoc("Omid", 27),
	}
	_, err := mgm.Coll(&Doc{}).InsertMany(mgm.Ctx(), docs)

	util.PanicErr(err)
}

func findDoc(t *testing.T) *Doc {
	found := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(found).FindOne(mgm.Ctx(), bson.M{}).Decode(found))

	return found
}

type Doc struct {
	mgm.DefaultModel `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func NewDoc(name string, age int) *Doc {
	return &Doc{Name: name, Age: age}
}
