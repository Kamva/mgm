package crud

import (
	"github.com/Kamva/mgm"
	"github.com/Kamva/mgm/builder"
	"github.com/Kamva/mgm/field"
	. "go.mongodb.org/mongo-driver/bson"
)

func lookup() error {

	author := newAuthor("Mehran")
	if err := mgm.Coll(author).Create(author); err != nil {
		return err
	}

	book := newBook("Test", 124, author.Id)
	booksColl := mgm.Coll(book)

	if err := booksColl.Create(book); err != nil {
		return err
	}

	_,err:=booksColl.Aggregate(mgm.Ctx(),A{
		builder.S{builder.Lookup(booksColl.Name(),"author_id",field.Id,"author")},
	})

	if err != nil {
		return err
	}

	return nil

}
