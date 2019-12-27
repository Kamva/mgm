package mgm

import (
	"github.com/Kamva/mgm/builder"
	"github.com/Kamva/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection performs operations on models and given Mongodb collection
type Collection struct {
	*mongo.Collection
}

// FindByID method find a doc and decode it to model, otherwise return error.
// id field can be any value that if passed to `PrepareID` method, it return
// valid id(e.g string,bson.ObjectId).
func (coll *Collection) FindByID(id interface{}, model Model) error {
	id, err := model.PrepareID(id)

	if err != nil {
		return err
	}

	return first(coll, bson.M{field.Id: id}, model)
}

// First method search and return first document of search result.
func (coll *Collection) First(filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return first(coll, filter, model, opts...)
}

// Create method insert new model into database.
func (coll *Collection) Create(model Model) error {
	return create(coll, model)
}

// Update function update save changed model into database.
// On call to this method also mgm call to model's updating,updated,
// saving,saved hooks.
func (coll *Collection) Update(model Model) error {
	return update(coll, model)
}

// Save method save model(insert,update).
func (coll *Collection) Save(model Model) error {
	if model.IsNew() {
		return create(coll, model)
	}

	return update(coll, model)
}

// Delete method delete model (doc) from collection.
// If you want to doing something on deleting some model
// use hooks, don't need to override this method.
func (coll *Collection) Delete(model Model) error {
	return del(coll, model)
}

// SimpleFind find and decode result to results.
func (coll *Collection) SimpleFind(results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	ctx := ctx()
	cur, err := coll.Find(ctx, filter, opts...)

	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

//--------------------------------
// Aggregation methods
//--------------------------------

// SimpleAggregate doing simple aggregation and decode aggregate result to the results.
// stages value can be Operator|bson.M
func (coll *Collection) SimpleAggregate(results interface{}, stages ...interface{}) error {
	cur, err := coll.SimpleAggregateCursor(stages...)
	if err != nil {
		return err
	}

	return cur.All(ctx(), results)
}

// SimpleAggregateCursor doing simple aggregation and return cursor.
func (coll *Collection) SimpleAggregateCursor(stages ...interface{}) (*mongo.Cursor, error) {
	pipeline := bson.A{}

	for _, stage := range stages {
		if operator, ok := stage.(builder.Operator); ok {
			pipeline = append(pipeline, builder.S(operator))
		} else {
			pipeline = append(pipeline, stage)
		}
	}

	return coll.Aggregate(ctx(), pipeline, nil)
}
