package mgm

// ModelCollection return model's collection.
func ModelCollection(m Model) *Collection {
	return m.Collection()
}