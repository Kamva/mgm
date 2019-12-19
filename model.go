package mgm

// By Implementing CollectionGetter interface for model's
// struct, you can returning model's custom collection.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

type Model interface {
	CollectionName() string

	// PrepareId convert id value if need, and then
	// return it.(e.g convert string to objectId)
	PrepareId(id interface{}) (interface{}, error)

	IsNew() bool
	GetId() interface{}
	SetId(id interface{})
}

type DefaultModel struct {
	IdField    `bson:",inline"`
	DateFields `bson:",inline"`
}

func (model *DefaultModel) Creating() error {
	return model.DateFields.Creating()
}

func (model *DefaultModel) Saving() error {
	return model.DateFields.Saving()
}
