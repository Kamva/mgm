package mgm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatingHook is called before saving a new model to the database
type CreatingHook interface {
	Creating(context.Context) error
}

// CreatedHook is called after a model has been created
type CreatedHook interface {
	Created(context.Context) error
}

// UpdatingHook is called before updating a model
type UpdatingHook interface {
	Updating(context.Context) error
}

// UpdatedHook is called after a model is updated
type UpdatedHook interface {
	Updated(ctx context.Context, result *mongo.UpdateResult) error
}

// SavingHook is called before a model (new or existing) is saved to the database.
type SavingHook interface {
	Saving(context.Context) error
}

// SavedHook is called after a model is saved to the database.
type SavedHook interface {
	Saved(context.Context) error
}

// DeletingHook is called before a model is deleted
type DeletingHook interface {
	Deleting(context.Context) error
}

// DeletedHook is called after a model is deleted
type DeletedHook interface {
	Deleted(ctx context.Context, result *mongo.DeleteResult) error
}

func callToBeforeCreateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(CreatingHook); ok {
		if err := hook.Creating(ctx); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeUpdateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(UpdatingHook); ok {
		if err := hook.Updating(ctx); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterCreateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(CreatedHook); ok {
		if err := hook.Created(ctx); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(ctx); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterUpdateHooks(ctx context.Context, updateResult *mongo.UpdateResult, model Model) error {
	if hook, ok := model.(UpdatedHook); ok {
		if err := hook.Updated(ctx, updateResult); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(ctx); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeDeleteHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(DeletingHook); ok {
		if err := hook.Deleting(ctx); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterDeleteHooks(ctx context.Context, deleteResult *mongo.DeleteResult, model Model) error {
	if hook, ok := model.(DeletedHook); ok {
		if err := hook.Deleted(ctx, deleteResult); err != nil {
			return err
		}
	}

	return nil
}
