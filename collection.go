package mgm

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	*mongo.Collection
}

// Find a doc and decode it to model, otherwise return error
func (coll *Collection) First(id interface{}, model Model) error {
	id, err := model.PrepareId(id)

	if err != nil {
		return err
	}

	return find(coll, id, model)
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
