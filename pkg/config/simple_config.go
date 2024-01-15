package config

import (
	"time"

	"github.com/adaptive-scale/dbchaos/pkg/runner"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const (
	MySQL     = "mysql"
	Postgres  = "postgres"
	SQLServer = "sqlserver"
)

type Randomize struct {
	Variable           string `yaml:"variable"`
	RangeFrom          int    `yaml:"rangeFrom,omitempty"`
	RangeTo            int    `yaml:"rangeTo,omitempty"`
	RandomString       bool   `yaml:"randomString,omitempty"`
	RandomAlphanumeric bool   `yaml:"randomAlphanumeric,omitempty"`
	maxLength          int    `yaml:"maxLength,omitempty"`
}

type SimpleConfiguration struct {
	DbType           string    `json:"db_type" yaml:"dbType"`
	ConnectionString string    `json:"connection_string" yaml:"connection,omitempty"`
	Query            string    `json:"query" yaml:"query"`
	ParallelRuns     int       `json:"parallel_runs" yaml:"parallelRuns,omitempty"`
	RunFor           string    `json:"run_for" yaml:"runFor,omitempty"`
	CoolOffTime      int       `json:"coolOffTime" yaml:"coolOffTime,omitempty"`
	Randomize        Randomize `json:"randomize" yaml:"randomize,omitempty"`
	//RequestPerSecond int64     `yaml:"requestPerSecond" json:"requestPerSecond"`
}

func ParseConfiguration(config []byte) *SimpleConfiguration {
	var configuration SimpleConfiguration
	if err := yaml.Unmarshal(config, &configuration); err != nil {
		return nil
	}
	return &configuration
}

func (s *SimpleConfiguration) Start() error {

	var d *gorm.DB
	var err error
	switch s.DbType {
	case MySQL:
		{
			d, err = gorm.Open(mysql.Open(s.ConnectionString), &gorm.Config{})
			if err != nil {
				return err
			}

		}
	case Postgres:
		{
			d, err = gorm.Open(postgres.Open(s.ConnectionString), &gorm.Config{})
			if err != nil {
				return err
			}
		}
	case SQLServer:
		{
			d, err = gorm.Open(sqlserver.Open(s.ConnectionString), &gorm.Config{})
			if err != nil {
				return err
			}
		}
	}

	sqlDB, _ := d.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if s.RunFor != "" {
		dur := runner.DurationRunner{
			RunFor:       s.RunFor,
			ParallelRuns: s.ParallelRuns,
			CoolOffTime:  s.CoolOffTime,
			DB:           d,
			//RequestPerSecond: s.RequestPerSecond,
			Query: s.Query,
		}
		return dur.Run()
	}

	return nil
}
