package mgm

import "go.mongodb.org/mongo-driver/mongo"

// CreatingHook is called before saving a new model to the database
type CreatingHook interface {
	Creating() error
}

// CreatedHook is called after a model has been created
type CreatedHook interface {
	Created() error
}

// UpdatingHook is called before updating a model
type UpdatingHook interface {
	Updating() error
}

// UpdatedHook is called after a model is updated
type UpdatedHook interface {
	Updated(result *mongo.UpdateResult) error
}

// SavingHook is called before a model (new or existing) is saved to the database.
type SavingHook interface {
	Saving() error
}

// SavedHook is called after a model is saved to the database.
type SavedHook interface {
	Saved() error
}

// DeletingHook is called before a model is deleted
type DeletingHook interface {
	Deleting() error
}

// DeletedHook is called after a model is deleted
type DeletedHook interface {
	Deleted(result *mongo.DeleteResult) error
}

func callToBeforeCreateHooks(model Model) error {
	if hook, ok := model.(CreatingHook); ok {
		if err := hook.Creating(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeUpdateHooks(model Model) error {
	if hook, ok := model.(UpdatingHook); ok {
		if err := hook.Updating(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterCreateHooks(model Model) error {
	if hook, ok := model.(CreatedHook); ok {
		if err := hook.Created(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterUpdateHooks(updateResult *mongo.UpdateResult, model Model) error {
	if hook, ok := model.(UpdatedHook); ok {
		if err := hook.Updated(updateResult); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeDeleteHooks(model Model) error {
	if hook, ok := model.(DeletingHook); ok {
		if err := hook.Deleting(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterDeleteHooks(deleteResult *mongo.DeleteResult, model Model) error {
	if hook, ok := model.(DeletedHook); ok {
		if err := hook.Deleted(deleteResult); err != nil {
			return err
		}
	}

	return nil
}
