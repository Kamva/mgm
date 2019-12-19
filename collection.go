package mgm

import (
	"github.com/Kamva/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	*mongo.Collection
}

// Find a doc and decode it to model, otherwise return error
func (coll *Collection) FindById(id interface{}, model Model) error {
	id, err := model.PrepareId(id)

	if err != nil {
		return err
	}

	return first(coll, bson.M{field.Id: id}, model)
}

func (coll *Collection) First(filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return first(coll, filter, model, opts...)
}

func (coll *Collection) Create(model Model) error {
	return create(coll, model)
}

func (coll *Collection) Update(model Model) error {
	return update(coll, model)
}

func (coll *Collection) Save(model Model) error {
	if model.IsNew() {
		return create(coll, model)
	}

	return update(coll, model)
}

func (coll *Collection) Delete(model Model) error {
	return del(coll, model)
}
