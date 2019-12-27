package crud

import "github.com/Kamva/mgm"

func crud() error {

	author := newAuthor("Mehran")
	if err := mgm.Coll(author).Create(author); err != nil {
		return err
	}

	book := newBook("Test", 124, author.Id)
	booksColl := mgm.Coll(book)

	if err := booksColl.Create(book); err != nil {
		return err
	}

	book.Name = "Moulin Rouge!"
	if err := booksColl.Save(book); err != nil {
		return err
	}

	if err := booksColl.Delete(book); err != nil {
		return err
	}

	return mgm.Coll(author).Delete(author)
}

func find() {
	// Get document's collection
	mgm.Coll(&book{})
}
