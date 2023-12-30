package runner

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
	"time"
)

type DurationRunner struct {
	RunFor           string
	Query            string
	ParallelRuns     int
	CoolOffTime      int
	DB               *gorm.DB
	RequestPerSecond int64
}

func (d *DurationRunner) Run() error {

	runID := uuid.New().String()

	t := time.Now()
	d1, err := time.ParseDuration(d.RunFor)
	if err != nil {
		return err
	}

	var v struct{}
	var totalQueryCount int64
	c := time.Duration(d.CoolOffTime) * time.Second
	for time.Since(t) <= d1 {
		if totalQueryCount%1000 == 0 {
			fmt.Println("["+runID+"] : total queries executed so far = ", totalQueryCount)
		}
		var wg sync.WaitGroup
		wg.Add(d.ParallelRuns)
		for i := 0; i < d.ParallelRuns; i++ {
			totalQueryCount++
			go func() {
				a := strings.Split(d.Query, ";")
				for _, a1 := range a {
					if strings.TrimSpace(a1) != "" {
						if err := d.DB.Raw(a1).Scan(&v).Error; err != nil {
							log.Println(err)
						}
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()

		time.Sleep(c)
	}
	fmt.Println("["+runID+"] finished running chaos. total queries executed = ", totalQueryCount)

	return nil
}
