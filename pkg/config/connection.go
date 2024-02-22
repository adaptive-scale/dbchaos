package config

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

type Connection struct {
	DbType           string `json:"db_type,omitempty" yaml:"dbType,omitempty"`
	ConnectionString string `json:"connection_string,omitempty" yaml:"connection,omitempty"`
	DbName           string `json:"db_name,omitempty" yaml:"dbName,omitempty"`
}

func (s Connection) NewClient() (*gorm.DB, error) {

	var d *gorm.DB
	var err error

	switch s.DbType {
	case MySQL:
		{
			d, err = gorm.Open(mysql.Open(s.ConnectionString), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			sqlDB, _ := d.DB()

			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			return d, nil

		}
	case Postgres:
		{
			d, err = gorm.Open(postgres.Open(s.ConnectionString), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			sqlDB, _ := d.DB()

			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			return d, nil

		}
	case SQLServer, "mssql":
		{
			d, err = gorm.Open(sqlserver.Open(s.ConnectionString), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			sqlDB, _ := d.DB()

			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			return d, nil
		}
	default:
		{
			return nil, errors.New("db is not supported")
		}
	}

}
