package config

import (
	"context"
	"log"
	"time"

	"github.com/adaptive-scale/dbchaos/pkg/runner"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

const (
	MySQL     = "mysql"
	Postgres  = "postgres"
	SQLServer = "sqlserver"
	MongoDB   = "mongodb"
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
	Connection
	Collection      string        `json:"collection" yaml:"collection"`
	Query           string        `json:"query" yaml:"query"`
	ParallelRuns    int           `json:"parallel_runs" yaml:"parallelRuns,omitempty"`
	RunFor          string        `json:"run_for" yaml:"runFor,omitempty"`
	CoolOffTime     int           `json:"coolOffTime" yaml:"coolOffTime,omitempty"`
	Randomize       Randomize     `json:"randomize" yaml:"randomize,omitempty"`
	QueryType       string        `json:"query_type" yaml:"queryType"`             // Applies to MongoDB Only
	SortQuery       string        `json:"sort_query" yaml:"sortQuery"`             // Applies to MongoDB Only
	SkipNumber      int           `json:"skip_number" yaml:"skipNumber"`           // Applies to MongoDB Only
	LimitNumber     int           `json:"limit_number" yaml:"limitNumber"`         // Applies to MongoDB Only
	ProjectionQuery string        `json:"projection_query" yaml:"projectionQuery"` // Applies to MongoDB Only
	Docs            []interface{} `json:"docs" yaml:"docs"`                        // Applies to NoSQL Databases Only
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
	case MySQL, Postgres, SQLServer:
		{
			d, err = s.NewClient()
			if err != nil {
				return err
			}
		}
	case MongoDB:
		{
			ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
			clientOptions := options.Client().ApplyURI(s.ConnectionString)
			clientOptions.SetMaxPoolSize(100)
			client, err := mongo.Connect(ctx, clientOptions)
			defer cancel()
			if err != nil {
				log.Fatal(err)
			}
			if s.RunFor != "" {
				dur := runner.DurationRunner{
					RunFor:       s.RunFor,
					ParallelRuns: s.ParallelRuns,
					CoolOffTime:  s.CoolOffTime,
					MongoDB:      client,
					DbType:       s.DbType,
					DbName:       s.DbName,
					//RequestPerSecond: s.RequestPerSecond,
					Query: s.Query,
				}
				return dur.Run()
			}
		}
	}

	if s.RunFor != "" && s.DbType != MongoDB {
		dur := runner.DurationRunner{
			RunFor:       s.RunFor,
			ParallelRuns: s.ParallelRuns,
			CoolOffTime:  s.CoolOffTime,
			DB:           d,
			DbType:       s.DbType,
			//RequestPerSecond: s.RequestPerSecond,
			Query: s.Query,
		}
		return dur.Run()
	}

	return nil
}
