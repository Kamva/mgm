package mgm

// CollectionGetter interface contains a method to return
// a model's custom collection.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

// CollectionNameGetter interface contains a method to return
// the collection name of a model.
type CollectionNameGetter interface {
	// CollectionName method return model collection's name.
	CollectionName() string
}

// Model interface contains base methods that must be implemented by
// each model. If you're using the `DefaultModel` struct in your model,
// you don't need to implement any of these methods.
type Model interface {
	// PrepareID converts the id value if needed, then
	// returns it (e.g convert string to objectId).
	PrepareID(id interface{}) (interface{}, error)

	GetID() interface{}
	SetID(id interface{})
}

// DefaultModel struct contains a model's default fields.
type DefaultModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

// Creating function calls the inner fields' defined hooks
// TODO: get context as param in the next version (4).
func (model *DefaultModel) Creating() error {
	return model.DateFields.Creating()
}

// Saving function calls the inner fields' defined hooks
// TODO: get context as param the next version(4).
func (model *DefaultModel) Saving() error {
	return model.DateFields.Saving()
}
