/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"github.com/jinzhu/gorm"
	"luban/pkg/database"
)

var config *Config

// InitConfig Initialize all used configurations
func InitConfig(cfg *Config) error {
	config = cfg
	err := initDatabase(&cfg.Database)
	return err
}

var db *gorm.DB

// InitDatabase Initialize database
func initDatabase(cfg *database.Config) error {
	var err error
	db, err = database.New(cfg)
	return err
}
