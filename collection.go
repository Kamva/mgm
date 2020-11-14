package mgm

import (
	"context"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/field"
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
	return coll.FindByIDWithCtx(ctx(), id, model)
}

// FindByIDWithCtx method find a doc and decode it to model, otherwise return error.
// id field can be any value that if passed to `PrepareID` method, it return
// valid id(e.g string,bson.ObjectId).
func (coll *Collection) FindByIDWithCtx(ctx context.Context, id interface{}, model Model) error {
	id, err := model.PrepareID(id)

	if err != nil {
		return err
	}

	return first(ctx, coll, bson.M{field.ID: id}, model)
}

// First method search and return first document of search result.
func (coll *Collection) First(filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return coll.FirstWithCtx(ctx(), filter, model, opts...)
}

// FirstWithCtx method search and return first document of search result.
func (coll *Collection) FirstWithCtx(ctx context.Context, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return first(ctx, coll, filter, model, opts...)
}

// Create method insert new model into database.
func (coll *Collection) Create(model Model, opts ...*options.InsertOneOptions) error {
	return coll.CreateWithCtx(ctx(), model, opts...)
}

// CreateWithCtx method insert new model into database.
func (coll *Collection) CreateWithCtx(ctx context.Context, model Model, opts ...*options.InsertOneOptions) error {
	return create(ctx, coll, model, opts...)
}

// Update function update save changed model into database.
// On call to this method also mgm call to model's updating,updated,
// saving,saved hooks.
func (coll *Collection) Update(model Model, opts ...*options.UpdateOptions) error {
	return coll.UpdateWithCtx(ctx(), model, opts...)
}

// UpdateWithCtx function update save changed model into database.
// On call to this method also mgm call to model's updating,updated,
// saving,saved hooks.
func (coll *Collection) UpdateWithCtx(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	return update(ctx, coll, model, opts...)
}

// Delete method delete model (doc) from collection.
// If you want to doing something on deleting some model
// use hooks, don't need to override this method.
func (coll *Collection) Delete(model Model) error {
	return del(ctx(), coll, model)
}

// DeleteWithCtx method delete model (doc) from collection.
// If you want to doing something on deleting some model
// use hooks, don't need to override this method.
func (coll *Collection) DeleteWithCtx(ctx context.Context, model Model) error {
	return del(ctx, coll, model)
}

// SimpleFind find and decode result to results.
func (coll *Collection) SimpleFind(results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	return coll.SimpleFindWithCtx(ctx(), results, filter, opts...)
}

// SimpleFindWithCtx find and decode result to results.
func (coll *Collection) SimpleFindWithCtx(ctx context.Context, results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := coll.Find(ctx, filter, opts...)

	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

//--------------------------------
// Aggregation methods
//--------------------------------

// SimpleAggregateFirst does simple aggregation and decode first aggregate result to the provided result param.
// stages value can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not get context.
// So you should use the regular aggregation method in transactions.
func (coll *Collection) SimpleAggregateFirst(result interface{}, stages ...interface{}) (bool, error) {
	cur, err := coll.SimpleAggregateCursor(stages...)
	if err != nil {
		return false, err
	}
	if cur.Next(ctx()) {
		return true, cur.Decode(result)
	}
	return false, nil
}

// SimpleAggregate does simple aggregation and decode aggregate result to the results.
// stages value can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not get context.
//So you should use the regular aggregation method in transactions.
func (coll *Collection) SimpleAggregate(results interface{}, stages ...interface{}) error {
	cur, err := coll.SimpleAggregateCursor(stages...)
	if err != nil {
		return err
	}

	return cur.All(ctx(), results)
}

// SimpleAggregateCursor doing simple aggregation and return cursor.
// Note: you can not use this method in a transaction because it does not get context.
// So you should use the regular aggregation method in transactions.
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
