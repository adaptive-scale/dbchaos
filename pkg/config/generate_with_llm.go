package config

import (
	"errors"
	"fmt"
	"github.com/adaptive-scale/dbchaos/pkg/generatewithllm"
	openai_dbchaos "github.com/adaptive-scale/dbchaos/pkg/openai"
	"log"
	"os"
	"strings"
	"time"
)

type SchemaGenerationWithLLM struct {
	Connection       Connection `yaml:"connection,omitempty"`
	Provider         string     `yaml:"provider,omitempty"`
	Model            string     `yaml:"model,omitempty"`
	SchemaType       string     `yaml:"schema_type,omitempty"`
	InsertIterations int        `yaml:"insert_iterations,omitempty"`
	DryRun           bool       `yaml:"dry_run,omitempty"`
	PersistSchema    string     `yaml:"persist_schema,omitempty"`
}

const (
	PromptTemplateForGenerateSchema = "Generate Schema in SQL for %v database. Give me only SQL Commands No text. Make tables with foreign key references are at the end."
	PromptTemplateForInsertData     = "Also generate the 100 SQL commands insert randomized data into the %v. Give me only SQL Commands No text."
)

func (s *SchemaGenerationWithLLM) GenerateSchema(apiToken string) error {

	totalWords := strings.Split(s.SchemaType, " ")
	if len(totalWords) > 1 {
		return errors.New("schema type can only be one word")
	}

	log.Println("starting schema generation with LLM")

	switch s.Provider {
	case "openai":
		return s.generateDataWithOpenAI(apiToken)
	case "llama-7b":
		return errors.New("llama-7b is not supported")
	case "Mixtral-8x7B":
		return errors.New("Mixtral-8x7B is not supported")
	default:
		return s.generateDataWithOpenAI(apiToken)
	}
}

func (s *SchemaGenerationWithLLM) generateDataWithOpenAI(apiToken string) error {
	o := openai_dbchaos.OpenAI{
		Model:  s.Model,
		APIkey: apiToken,
	}

	log.Println("generating schema")

	schemaValue := generatewithllm.KnownSchema[strings.ToLower(s.SchemaType)]

	if schemaValue == "" {
		v, err := o.Prompt(fmt.Sprintf(PromptTemplateForGenerateSchema, s.SchemaType))
		if err != nil {
			return err
		}
		schemaValue = v
	}

	var dataVal string

	if s.InsertIterations == 0 {
		s.InsertIterations = 1
	}

	log.Println("generating data")

	for i := 0; i < s.InsertIterations; i++ {
		// Insert data
		data, err := o.Prompt(fmt.Sprintf(PromptTemplateForInsertData, schemaValue))
		if err != nil {
			return err
		}
		dataVal = dataVal + "\n--insert data\n" + data

		time.Sleep(1 * time.Second)
	}

	queries := schemaValue + "\n--insert data\n" + dataVal

	if s.PersistSchema != "" {

		log.Println("saving commands to file to " + s.PersistSchema)

		err := os.WriteFile(s.PersistSchema, []byte(queries), 0644)
		if err != nil {
			return err
		}
	}

	if s.DryRun {
		fmt.Println(queries)
	} else {
		d, err := s.Connection.NewClient()
		if err != nil {
			return err
		}

		log.Println("executing commands")

		fmt.Println(queries)

		if err := d.Exec(queries).Error; err != nil {
			return err
		}
	}

	return nil
}
