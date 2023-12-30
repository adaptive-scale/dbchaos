package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"sync"
)

type Scenario struct {
	Scenarios        []SimpleConfiguration `json:"scenarios" yaml:"scenarios"`
	DbType           string                `json:"db_type" yaml:"dbType"`
	ConnectionString string                `json:"connection_string" yaml:"connection"`
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

		go func(a SimpleConfiguration) {
			if err := a.Start(); err != nil {
				log.Println(err)
				wg.Done()
				return
			}
		}(a)
	}

	wg.Wait()

}
