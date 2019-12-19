package mgm

import (
	"go.mongodb.org/mongo-driver/bson"
)

func create(c *Collection, model Model) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx(), model)

	if err != nil {
		return err
	}

	// Set new id
	model.SetId(res.InsertedID)

	return callToAfterCreateHooks(model)
}

func find(c *Collection, id interface{}, model Model) error {
	return c.FindOne(ctx(), bson.M{"_id": id}).Decode(model)
}

func update(c *Collection, model Model) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(model); err != nil {
		return err
	}

	res, err := c.UpdateOne(ctx(), bson.M{"_id": model.GetId()}, bson.M{"$set": model})

	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(res, model)
}

func del(c *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(model); err != nil {
		return err
	}
	res, err := c.DeleteOne(ctx(), bson.M{"_id": model.GetId()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(res, model)
}
