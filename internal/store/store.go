package store

import (
	"casher-server/internal/config"
	"context"

	"github.com/w6xian/sqlm"
	"go.uber.org/zap"
)

type Store struct {
	profile *config.Profile
	driver  Driver
	lager   *zap.Logger
}

func New(driver Driver, opt *config.Profile, lager *zap.Logger) (*Store, error) {
	store := &Store{
		profile: opt,
		driver:  driver,
		lager:   lager,
	}
	return store, nil
}

func (s *Store) GetDriver() Driver {
	return s.driver
}

func (s *Store) GetConnect(ctx context.Context) context.Context {
	return s.driver.GetConnect(ctx)
}

func (s *Store) CloseConnect(ctx context.Context) error {
	return s.driver.CloseConnect(ctx)
}

func (s *Store) Close() error {
	// Stop all cache cleanup goroutines
	return s.driver.Close()
}

func (s *Store) GetLink(ctx context.Context) sqlm.ITable {
	return s.driver.GetLink(ctx)
}

func (s *Store) Action(ctx context.Context, f func(tx sqlm.ITable, args ...any) (int64, error)) (int64, error) {
	link := s.driver.GetAction(ctx)
	defer link.Close()
	return link.Action(f)
}

func (v *Store) DbConnectWithClose(ctx context.Context) (context.Context, func()) {
	return v.driver.GetConnect(ctx), func() {
		v.driver.CloseConnect(ctx)
	}
}
