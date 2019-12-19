package mgm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IdField struct {
	Id primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}

type DateFields struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (f *IdField) PrepareId(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

func (f *IdField) IsNew() bool {
	return f.GetId() == primitive.ObjectID{}
}

func (f *IdField) GetId() interface{} {
	return f.Id
}

func (f *IdField) SetId(id interface{}) {
	f.Id = id.(primitive.ObjectID)
}

//--------------------------------
// DateField methods
//--------------------------------

func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}
