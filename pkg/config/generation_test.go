package config

import (
	"log"
	"testing"
)

// SchemaType is a type of schema
func TestSchemaGeneration(t *testing.T) {
	var schemaConfig SchemaGeneration

	schemaConfig.Connection.DbType = "postgres"
	schemaConfig.Connection.DbName = "test-sql"

	//password := url.QueryEscape(`adaptive-3jvw:G!q6BVhjZz^0`)
	//schemaConfig.ConnectionString = `sqlserver://` + `adaptive-3jvw:G!q6BVhjZz^0` + `@127.0.0.1:11433?database=pubs`

	schemaConfig.Connection.ConnectionString = `host=localhost user=postgres password=mysecretpassword dbname=test_generation port=5432 sslmode=disable`

	schemaConfig.Schema.NumberOfSchema = 1
	schemaConfig.Schema.GenerateTables = true
	schemaConfig.Schema.Language = "en"
	schemaConfig.DryRun = false
	schemaConfig.Tables.NumberOfTables = 20
	schemaConfig.Tables.MinColumns = 1
	schemaConfig.Tables.MaxColumns = 5

	schemaConfig.Tables.PopulateTable = true
	schemaConfig.Rows.MinRows = 1
	schemaConfig.Rows.MaxRows = 10

	err := schemaConfig.GenerateSchema()
	if err != nil {
		log.Fatal(err)
	}
}
