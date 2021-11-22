package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField struct contains a model's ID field.
type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// DateFields struct contains the `created_at` and `updated_at`
// fields that autofill when inserting or updating a model.
type DateFields struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type VersionField struct {
	Version int `json:"version" bson:"version"`
}

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID
func (f *IDField) GetID() interface{} {
	return f.ID
}

// SetID sets the value of a model's ID field.
func (f *IDField) SetID(id interface{}) {
	f.ID = id.(primitive.ObjectID)
}

// GetVersion returns the model version field
func (f *VersionField) GetVersion() interface{} {
	return f.Version
}

// GetVersionFieldName returns the field name holding the version field (has to match the bson tag)
func (f *VersionField) GetVersionFieldName() string {
	return "version"
}

// SetVersion returns the model version field
func (f *VersionField) IncrementVersion() {
	f.Version++
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook is used here to set the `created_at` field
// value when inserting a new model into the database.
// TODO: get context as param the next version(4).
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook is used here to set the `updated_at` field
// value when creating or updateing a model.
// TODO: get context as param the next version(4).
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}
