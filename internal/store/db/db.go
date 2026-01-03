package db

import (
	"casher-server/internal/config"
	"casher-server/internal/store"
	"casher-server/internal/store/db/mysql"

	"github.com/pkg/errors"
)

// NewDBDriver creates new db driver based on profile.
func NewDBDriver(profile *config.Profile) (store.Driver, error) {
	var driver store.Driver
	var err error
	switch profile.Store.Protocol {
	case "mysql":
		driver, err = mysql.NewDB(profile)
	case "sqlite":
		return nil, errors.New("sqlite driver not implemented")
	default:
		return nil, errors.New("unknown db driver")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db driver")
	}
	// fmt.Println("NewDBDriver driver=", driver)
	// fmt.Println("NewDBDriver driver.Protocol=", profile.Store.Protocol)
	return driver, nil
}
