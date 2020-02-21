package mgm

// CollectionGetter interface contain method to return
// model's custom collection.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

// CollectionNameGetter interface contain method to return
// collection name of model.
type CollectionNameGetter interface {
	// CollectionName method return model collection's name.
	CollectionName() string
}

// Model interface is base method that must implement by
// each model, If you're using `DefaultModel` struct in your model,
// don't need to implement any of those method.
type Model interface {
	// PrepareID convert id value if need, and then
	// return it.(e.g convert string to objectId)
	PrepareID(id interface{}) (interface{}, error)

	GetID() interface{}
	SetID(id interface{})
}

// DefaultModel struct contain model's default fields.
type DefaultModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

// Creating function call to it's inner fields defined hooks
func (model *DefaultModel) Creating() error {
	return model.DateFields.Creating()
}

// Saving function call to it's inner fields defined hooks
func (model *DefaultModel) Saving() error {
	return model.DateFields.Saving()
}
