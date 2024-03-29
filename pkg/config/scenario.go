package config

import (
	"log"
	"sync"

	"gopkg.in/yaml.v3"
)

type Scenario struct {
	Connection                            // Applies to NoSQL Databases Only
	Scenarios       []SimpleConfiguration `json:"scenarios" yaml:"scenarios"`
	Collection      string                `json:"collection" yaml:"collection"`            // Applies to NoSQL Databases Only
	QueryType       string                `json:"query_type" yaml:"queryType"`             // Applies to MongoDB Only
	SortQuery       string                `json:"sort_query" yaml:"sortQuery"`             // Applies to MongoDB Only
	SkipNumber      int                   `json:"skip_number" yaml:"skipNumber"`           // Applies to MongoDB Only
	LimitNumber     int                   `json:"limit_number" yaml:"limitNumber"`         // Applies to MongoDB Only
	ProjectionQuery string                `json:"projection_query" yaml:"projectionQuery"` // Applies to MongoDB Only
	Docs            []interface{}         `json:"docs" yaml:"docs"`                        // Applies to NoSQL Databases Only
}

func ParseScenario(config []byte) *Scenario {
	var configuration Scenario
	if err := yaml.Unmarshal(config, &configuration); err != nil {
		return nil
	}
	return &configuration
}

func (s *Scenario) Start() {

	var wg sync.WaitGroup
	wg.Add(len(s.Scenarios))
	for _, a := range s.Scenarios {

		if a.DbType == "" {
			a.DbType = s.DbType
		}

		if a.ConnectionString == "" {
			a.ConnectionString = s.ConnectionString
		}

		a.DbName = s.DbName

		go func(a SimpleConfiguration) {
			defer wg.Done()
			if err := a.Start(); err != nil {
				log.Println(err)
				return
			}
		}(a)
	}

	wg.Wait()

}
