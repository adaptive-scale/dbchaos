package config

import "testing"

func TestGenerateWithLLM(t *testing.T) {

	var s SchemaGenerationWithLLM

	s.Connection.DbType = "postgres"
	s.Connection.DbName = "test-sql"
	s.Connection.ConnectionString = `host=localhost user=postgres password=mysecretpassword dbname=test_generation port=5432 sslmode=disable`

	s.Provider = "openai"
	s.Model = "gpt-3.5-turbo"
	s.SchemaType = "Webshop"

	s.InsertIterations = 1

	if err := s.GenerateSchema(""); err != nil {
		t.Errorf("Error: %v", err)
	}
}
