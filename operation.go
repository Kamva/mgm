package mgm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(ctx context.Context, c *Collection, model Model, opts ...*options.InsertOneOptions) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx, model, opts...)

	if err != nil {
		return err
	}

	// Set new id
	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(model)
}

func first(ctx context.Context, c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return c.FindOne(ctx, filter, opts...).Decode(model)
}

func update(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx, bson.M{model.PKField(): model.GetID()}, bson.M{"$set": model}, opts...)

	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(res, model)
}

func del(ctx context.Context, c *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(model); err != nil {
		return err
	}
	res, err := c.DeleteOne(ctx, bson.M{model.PKField(): model.GetID()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(res, model)
}
