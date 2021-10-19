package mgm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatingHook is called before saving a new model to the database
// Deprecated: please use CreatingHookWithCtx
type CreatingHook interface {
	Creating() error
}

// CreatingHookWithCtx is called before saving a new model to the database
type CreatingHookWithCtx interface {
	Creating(context.Context) error
}

// CreatedHook is called after a model has been created
// Deprecated: Please use CreatedHookWithCtx
type CreatedHook interface {
	Created() error
}

// CreatedHookWithCtx is called after a model has been created
type CreatedHookWithCtx interface {
	Created(context.Context) error
}

// UpdatingHook is called before updating a model
// Deprecated: Please use UpdatingHookWithCtx
type UpdatingHook interface {
	Updating() error
}

// UpdatingHookWithCtx is called before updating a model
type UpdatingHookWithCtx interface {
	Updating(context.Context) error
}

// UpdatedHook is called after a model is updated
// Deprecated: Please use UpdatedHookWithCtx
type UpdatedHook interface {
	// Deprecated:
	Updated(result *mongo.UpdateResult) error
}

// UpdatedHookWithCtx is called after a model is updated
type UpdatedHookWithCtx interface {
	Updated(ctx context.Context, result *mongo.UpdateResult) error
}

// SavingHook is called before a model (new or existing) is saved to the database.
// Deprecated: Please use SavingHookWithCtx
type SavingHook interface {
	Saving() error
}

// SavingHookWithCtx is called before a model (new or existing) is saved to the database.
type SavingHookWithCtx interface {
	Saving(context.Context) error
}

// SavedHook is called after a model is saved to the database.
// Deprecated: Please use SavedHookWithCtx
type SavedHook interface {
	Saved() error
}

// SavedHookWithCtx is called after a model is saved to the database.
type SavedHookWithCtx interface {
	Saved(context.Context) error
}

// DeletingHook is called before a model is deleted
// Deprecated: Please use DeletingHookWithCtx
type DeletingHook interface {
	Deleting() error
}

// DeletingHookWithCtx is called before a model is deleted
type DeletingHookWithCtx interface {
	Deleting(context.Context) error
}

// DeletedHook is called after a model is deleted
// Deprecated: Please use DeletedHookWithCtx
type DeletedHook interface {
	Deleted(result *mongo.DeleteResult) error
}

// DeletedHookWithCtx is called after a model is deleted
type DeletedHookWithCtx interface {
	Deleted(ctx context.Context, result *mongo.DeleteResult) error
}

func callToBeforeCreateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(CreatingHookWithCtx); ok {
		if err := hook.Creating(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(CreatingHook); ok {
		if err := hook.Creating(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHookWithCtx); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeUpdateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(UpdatingHookWithCtx); ok {
		if err := hook.Updating(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(UpdatingHook); ok {
		if err := hook.Updating(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavingHookWithCtx); ok {
		if err := hook.Saving(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavingHook); ok {
		if err := hook.Saving(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterCreateHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(CreatedHookWithCtx); ok {
		if err := hook.Created(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(CreatedHook); ok {
		if err := hook.Created(); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHookWithCtx); ok {
		if err := hook.Saved(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterUpdateHooks(ctx context.Context, updateResult *mongo.UpdateResult, model Model) error {
	if hook, ok := model.(UpdatedHookWithCtx); ok {
		if err := hook.Updated(ctx, updateResult); err != nil {
			return err
		}
	} else if hook, ok := model.(UpdatedHook); ok {
		if err := hook.Updated(updateResult); err != nil {
			return err
		}
	}

	if hook, ok := model.(SavedHookWithCtx); ok {
		if err := hook.Saved(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(SavedHook); ok {
		if err := hook.Saved(); err != nil {
			return err
		}
	}

	return nil
}

func callToBeforeDeleteHooks(ctx context.Context, model Model) error {
	if hook, ok := model.(DeletingHookWithCtx); ok {
		if err := hook.Deleting(ctx); err != nil {
			return err
		}
	} else if hook, ok := model.(DeletingHook); ok {
		if err := hook.Deleting(); err != nil {
			return err
		}
	}

	return nil
}

func callToAfterDeleteHooks(ctx context.Context, deleteResult *mongo.DeleteResult, model Model) error {
	if hook, ok := model.(DeletedHookWithCtx); ok {
		if err := hook.Deleted(ctx, deleteResult); err != nil {
			return err
		}
	} else if hook, ok := model.(DeletedHook); ok {
		if err := hook.Deleted(deleteResult); err != nil {
			return err
		}
	}

	return nil
}
