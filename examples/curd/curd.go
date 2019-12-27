package curd

import "github.com/Kamva/mgm"

func CRUD() error {
	book := newBook("test", 124)
	coll := mgm.Coll(book)


	if err := coll.Create(book); err != nil {
		return err
	}

	book.Name = "Moulin Rouge!"
	if err := coll.Save(book); err != nil {
		return err
	}

	return coll.Delete(book)
}

func find(){
	// Get document's collection
	mgm.Coll(&book{})
}