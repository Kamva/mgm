package crud

import (
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Pages            int                `json:"pages" bson:"pages"`
	AuthorId         primitive.ObjectID `json:"author_id" bson:"author_id"`
}

func newBook(name string, pages int, authId primitive.ObjectID) *book {
	return &book{
		Name:     name,
		Pages:    pages,
		AuthorId: authId,
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
