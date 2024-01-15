package runner

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type DurationRunner struct {
	RunFor           string
	Query            string
	ParallelRuns     int
	CoolOffTime      int
	DB               *gorm.DB
	MongoDB          *mongo.Client
	RequestPerSecond int64
	DbType           string
	DbName           string        // NoSQL Databases Only
	Collection       string        // NoSQL Databases Only
	QueryType        string        // Applies to MongoDB Only
	SortQuery        string        // Applies to MongoDB Only
	SkipNumber       int           // Applies to MongoDB Only
	LimitNumber      int           // Applies to MongoDB Only
	ProjectionQuery  string        // Applies to MongoDB Only
	Docs             []interface{} // Applies to NoSQL Databases Only
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
						if d.DbType == "mongodb" {
							var filter interface{}
							if d.Query != "" {
								err := bson.UnmarshalExtJSON([]byte(d.Query), true, &filter)
								if err != nil {
									log.Println(err)
								}
							}
							db := d.MongoDB.Database(d.DbName).Collection(d.Collection)
							switch d.QueryType {
							case "find":
								{
									_, err := db.Find(context.TODO(), filter)
									if err != nil {
										log.Println(err)
									}
								}
							case "findone":
								{
									_ = db.FindOne(context.TODO(), filter)
								}
							case "insertmany":
								{
									_, err := db.InsertMany(context.TODO(), d.Docs)
									if err != nil {
										log.Println(err)
									}
								}
							}
						} else {
							if err := d.DB.Raw(a1).Scan(&v).Error; err != nil {
								log.Println(err)
							}
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
