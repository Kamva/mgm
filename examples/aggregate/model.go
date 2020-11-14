package aggregate

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Pages            int                `json:"pages" bson:"pages"`
	AuthorID         primitive.ObjectID `json:"author_id" bson:"author_id"`
}

func newBook(name string, pages int, authID primitive.ObjectID) *book {
	return &book{
		Name:     name,
		Pages:    pages,
		AuthorID: authID,
	}
}

type author struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func newAuthor(name string) *author {
	return &author{
		Name: name,
	}
}
