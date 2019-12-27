package aggregate

import (
	"github.com/Kamva/mgm"
	"github.com/Kamva/mgm/builder"
	"github.com/Kamva/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
)

func seed() {
	author := newAuthor("Mehran")
	_ = mgm.Coll(author).Create(author)

	book := newBook("Test", 124, author.ID)
	_ = mgm.Coll(book).Create(book)

}

func delSeededData() {
	_, _ = mgm.Coll(&book{}).DeleteMany(nil, bson.M{})
	_, _ = mgm.Coll(&author{}).DeleteMany(nil, bson.M{})
}

func lookup() error {
	seed()

	defer delSeededData()

	// Author model's collection
	authorColl := mgm.Coll(&author{})

	pipeline := bson.A{
		builder.S(builder.Lookup(authorColl.Name(), "author_id", field.ID, "author")),
	}

	cur, err := mgm.Coll(&book{}).Aggregate(mgm.Ctx(), pipeline)

	if err != nil {
		return err
	}

	defer cur.Close(nil)

	for cur.Next(nil) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return err
		}

		// do something with result....
		//fmt.Printf("%+v\n", result)
	}

	return nil
}
