package mgm

import (
	"context"
	"fmt"
	"time"

	"github.com/kamva/mgm/v3/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(ctx context.Context, c *Collection, model Model, opts ...*options.InsertOneOptions) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(ctx, model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx, model, opts...)

	if err != nil {
		return err
	}

	// Set new id
	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(ctx, model)
}

func first(ctx context.Context, c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return c.FindOne(ctx, filter, opts...).Decode(model)
}

func update(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(ctx, model); err != nil {
		return err
	}

	query := bson.M{field.ID: model.GetID()}
	modelVersionable, isVersionable := model.(Versionable)
	var currentVersion interface{}
	if isVersionable {
		currentVersion = modelVersionable.GetVersion()

		//Handle adding versionning for documents that were created without it
		//In this case, the version field would be of zero value and not exist in the DB
		//Add $or exists false in the query condition for this case
		isCurrentVersionZero := false
		switch c := currentVersion.(type) {
		case string:
			isCurrentVersionZero = c == ""
		case int:
			isCurrentVersionZero = c == 0
		case time.Time:
			isCurrentVersionZero = c.IsZero()
		}

		if isCurrentVersionZero {
			query["$or"] = bson.A{bson.M{modelVersionable.GetVersionFieldName(): currentVersion}, bson.M{modelVersionable.GetVersionFieldName(): bson.M{"$exists": false}}}
		} else {
			query[modelVersionable.GetVersionFieldName()] = currentVersion
		}

		modelVersionable.IncrementVersion()
	}

	res, err := c.UpdateOne(ctx, query, bson.M{"$set": model}, opts...)

	if err != nil {
		return err
	}

	if isVersionable && res.MatchedCount == 0 {
		return fmt.Errorf("versioning error : document %v %v with version %v could not be found", c.Name(), model.GetID(), currentVersion)
	}

	return callToAfterUpdateHooks(ctx, res, model)
}

func del(ctx context.Context, c *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(ctx, model); err != nil {
		return err
	}
	res, err := c.DeleteOne(ctx, bson.M{field.ID: model.GetID()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(ctx, res, model)
}
