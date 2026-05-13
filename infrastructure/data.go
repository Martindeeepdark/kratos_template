package infrastructure

import (
	"context"

	"github.com/Martindeeepdark/go-common/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"kratos_template/internal/conf"
)

type Data struct {
	db *gorm.DB
}

func NewData(cfg conf.DatabaseConfig) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(cfg.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		logs.Info("closing the data resources")
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	return &Data{db: db}, cleanup, nil
}

func (d *Data) DB() *gorm.DB {
	return d.db
}

func (d *Data) Ping(ctx context.Context) error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
