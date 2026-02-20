package aggregate

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Pages            int                `json:"pages" bson:"pages"`
	AuthorID         bson.ObjectID `json:"author_id" bson:"author_id"`
}

func newBook(name string, pages int, authID bson.ObjectID) *book {
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
