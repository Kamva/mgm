package mgm

// Coll return model's collection.
func Coll(m Model) *Collection {

	if CollectionGetter, ok := m.(CollectionGetter); ok {
		return CollectionGetter.Collection()
	}

	return CollectionByName(m.CollectionName())
}
