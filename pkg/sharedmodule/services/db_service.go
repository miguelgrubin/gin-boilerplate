package services

import (
	"slices"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBService interface {
	Connect() error
	Close() error
	GetDB() *gorm.DB
}

type DBServiceGorm struct {
	driver string
	url    string
	DB     *gorm.DB
}

var _ DBService = &DBServiceGorm{}

func NewDBServiceGorm(dc DatabaseConfig) *DBServiceGorm {
	validDrivers := []string{"sqlite3", "mysql", "postgres"}
	if !slices.Contains(validDrivers, dc.Driver) {
		panic("Invalid database driver: " + dc.Driver)
	}

	return &DBServiceGorm{
		driver: dc.Driver,
		url:    dc.Address,
		DB:     nil,
	}
}

func (d *DBServiceGorm) Connect() error {
	driverMap := map[string]func(u string) gorm.Dialector{
		"sqlite3":  sqlite.Open,
		"mysql":    mysql.Open,
		"postgres": postgres.Open,
	}

	db, err := gorm.Open(driverMap[d.driver](d.url), &gorm.Config{})
	if err != nil {
		return err
	}
	d.DB = db

	return nil
}

func (d *DBServiceGorm) GetDB() *gorm.DB {
	return d.DB
}

func (d *DBServiceGorm) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
