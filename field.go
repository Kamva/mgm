package mgm

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// IDField struct contain model's ID field.
type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// DateFields struct contain `created_at` and `updated_at`
// fields that autofill on insert/update model.
type DateFields struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// PrepareID method prepare id value to using it as id in filtering,...
// e.g convert hex-string id value to bson.ObjectId
func (f *IDField) PrepareID(id interface{}) (oid primitive.ObjectID, err error) {
	switch t := id.(type) {
	case string:
		oid, err = primitive.ObjectIDFromHex(t)
		return
	case primitive.ObjectID:
		err = nil
		oid = t
		return
	default:
		err = errors.New("unknown type")
		return
	}
}

// GetID method return model's id
func (f *IDField) GetID() primitive.ObjectID {
	return f.ID
}

// SetID set id value of model's id field.
func (f *IDField) SetID(id interface{}) error {
	switch t := id.(type) {
	case string:
		id, err := primitive.ObjectIDFromHex(t)
		if err != nil {
			return err
		}
		f.ID = id
	case primitive.ObjectID:
		f.ID = t
	default:
		return errors.New("unknown type")
	}
	return nil
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook used here to set `created_at` field
// value on inserting new model into database.
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook used here to set `updated_at` field value
// on create/update model.
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}
