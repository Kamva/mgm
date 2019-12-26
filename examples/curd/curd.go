package curd

import "github.com/Kamva/mgm"

func CRUD() error {
	coll := mgm.Coll(&book{})

	book := newBook("test", 124)

	if err := coll.Create(book); err != nil {
		return err
	}

	book.Name = "Moulin Rouge!"
	if err := coll.Save(book); err != nil {
		return err
	}

	return coll.Delete(book)
}
