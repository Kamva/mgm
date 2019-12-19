package mgm

type Model interface {
	// Collection method return collection
	Collection() *Collection

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
