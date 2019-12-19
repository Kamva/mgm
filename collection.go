package mgm

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	*mongo.Collection
}

// Find a doc and decode it to model, otherwise return error
func (c *Collection) Find(id interface{}, model Model) error {
	id, err := model.PrepareId(id)

	if err != nil {
		return err
	}

	return find(c, id, model)
}

func (c *Collection) Create(model Model) error {
	return create(c, model)
}

func (c *Collection) Update(model Model) error {
	return update(c, model)
}

func (c *Collection) Save(model Model) error {
	if model.IsNew() {
		return create(c, model)
	}

	return update(c, model)
}

func (c *Collection) Delete(model Model) error {
	return del(c, model)
}
