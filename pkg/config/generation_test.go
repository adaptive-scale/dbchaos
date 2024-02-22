package config

import (
	"log"
	"testing"
)

// SchemaType is a type of schema
func TestSchemaGeneration(t *testing.T) {
	var schemaConfig StaticSchemaGeneration

	schemaConfig.DbType = "postgres"
	schemaConfig.DbName = "test-sql"

	//password := url.QueryEscape(`adaptive-3jvw:G!q6BVhjZz^0`)
	//schemaConfig.ConnectionString = `sqlserver://` + `adaptive-3jvw:G!q6BVhjZz^0` + `@127.0.0.1:11433?database=pubs`

	schemaConfig.ConnectionString = `host=localhost user=postgres password=mysecretpassword dbname=test_generation port=5432 sslmode=disable`

	schemaConfig.NumberOfSchema = 1
	schemaConfig.GenerateTables = true
	schemaConfig.Language = "en"
	schemaConfig.DryRun = false
	schemaConfig.GenerateTables = true
	schemaConfig.NumberOfTables = 20
	schemaConfig.MinColumns = 1
	schemaConfig.MaxColumns = 5

	schemaConfig.PopulateTable = true
	schemaConfig.MinRows = 1
	schemaConfig.MaxRows = 10

	err := schemaConfig.GenerateSchema()
	if err != nil {
		log.Fatal(err)
	}
}
