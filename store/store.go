package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/usememos/memos/server/profile"
)

// Store provides database access to all raw objects.
type Store struct {
	Profile            *profile.Profile
	db                 *sql.DB
	systemSettingCache sync.Map // map[string]*systemSettingRaw
	userCache          sync.Map // map[int]*userRaw
	userSettingCache   sync.Map // map[string]*userSettingRaw
	memoCache          sync.Map // map[int]*memoRaw
	shortcutCache      sync.Map // map[int]*shortcutRaw
	idpCache           sync.Map // map[int]*identityProviderMessage
}

// New creates a new instance of Store.
func New(db *sql.DB, profile *profile.Profile) *Store {
	return &Store{
		Profile: profile,
		db:      db,
	}
}

func (s *Store) Vacuum(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return FormatError(err)
	}
	defer tx.Rollback()

	if err := s.vacuumImpl(ctx, tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return FormatError(err)
	}

	// Vacuum sqlite database file size after deleting resource.
	if _, err := s.db.Exec("VACUUM"); err != nil {
		return err
	}

	return nil
}

func (*Store) vacuumImpl(ctx context.Context, tx *sql.Tx) error {
	if err := vacuumMemo(ctx, tx); err != nil {
		return err
	}
	if err := vacuumResource(ctx, tx); err != nil {
		return err
	}
	if err := vacuumShortcut(ctx, tx); err != nil {
		return err
	}
	if err := vacuumUserSetting(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMemoOrganizer(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMemoResource(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMemoRelations(ctx, tx); err != nil {
		return err
	}
	if err := vacuumTag(ctx, tx); err != nil {
		// Prevent revive warning.
		return err
	}

	return nil
}
