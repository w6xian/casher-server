package store

import (
	"casher-server/internal/config"
	"casher-server/internal/errors"
	"casher-server/internal/i18n"
	"context"
	"fmt"

	"github.com/w6xian/sqlm"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type Store struct {
	profile  *config.Profile
	driver   Driver
	lager    *zap.Logger
	Language string
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

func (v *Store) L(key string, def string, fields ...i18n.Field) string {
	if v.Language == "" {
		v.Language = language.Chinese.String()
	}
	l := len(fields)
	if l == 0 {
		return i18n.T(v.Language, key, def)
	}

	data := i18n.D{}
	for _, f := range fields {
		data[f.Key] = f.Value()
	}
	for k, v := range data {
		fmt.Printf("%s=%s\n", k, v)
	}
	return i18n.TWithData(v.Language, key, def, data)
}

func (v *Store) Error(key string, def string, fields ...i18n.Field) error {
	err := errors.FromLang(v)
	err.New(key, def, fields...)
	return errors.New(err.Error())
}
