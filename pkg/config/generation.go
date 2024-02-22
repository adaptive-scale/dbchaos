package config

import (
	"fmt"
	"github.com/adaptive-scale/dbchaos/pkg/namesgenerator"
	"github.com/adaptive-scale/dbchaos/pkg/utils"
	funktors "github.com/adaptive-scale/funktors"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Generation struct {
	// Name of the generation
	Connection
	StaticSchemaGeneration
}

type StaticSchemaGeneration struct {
	NumberOfSchema int    `json:"number_of_schema,omitempty" yaml:"numberOfSchema,omitempty"`
	GenerateTables bool   `json:"generate_tables,omitempty" yaml:"generateTables,omitempty"`
	Language       string `json:"language,omitempty" yaml:"language,omitempty"`
	DryRun         bool   `json:"dry_run,omitempty" yaml:"dryRun,omitempty"`
	TableGeneration
}

type internalTable struct {
	TableName string
	Fields    []Field
}

func (i internalTable) String(n int) string {

	fields := funktors.Map(i.Fields, func(i int, f Field) string {
		return f.Name
	})

	f := "INSERT INTO " + i.TableName + " (" + strings.Join(fields, ",") + ") VALUES "

	inserts := []string{}
	for iv := 0; iv < n; iv++ {
		rows := funktors.Map(i.Fields, func(i int, f Field) string {
			return f.Type.value()
		})

		inserts = append(inserts, "("+strings.Join(rows, ",")+")")
	}

	return f + strings.Join(inserts, ",")
}

type Field struct {
	Name string
	Type TypeInterface
}

type TypeInterface interface {
	name() string
	value() string
}

type intval struct{}

func (i intval) name() string {
	return "int"
}

func (i intval) value() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(19)
	return strconv.Itoa(randomNumber)
}

type floatval struct{}

func (f floatval) name() string {
	return "float"
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func (f floatval) value() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatFloat(rand.Float64(), 'f', -1, 64)
}

type textval struct{}

func (t textval) name() string {
	return "TEXT"
}

func (t textval) value() string {
	return `'` + utils.RandomString(utils.CharsetAlphabet, 10) + `'`
}

type varchar struct{}

func (v varchar) name() string {
	return "VARCHAR(255)"
}

func (v varchar) value() string {
	return `'` + utils.RandomString(utils.CharsetAlphabet, 10) + `'`
}

func randomizeType() TypeInterface {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3)
	switch randomNumber {
	case 0:
		return intval{}
	case 1:
		return floatval{}
	case 2:
		return textval{}
	case 3:
		return varchar{}
	default:
		return intval{}
	}
}

func (g StaticSchemaGeneration) GenerateSchema() error {

	totalSchema := g.NumberOfSchema

	var schemas []string

	for i := 0; i < totalSchema; i++ {
		name := namesgenerator.GetRandomName(2)
		schemas = append(schemas, "CREATE SCHEMA "+name)
	}

	fmt.Println("Schema Generation Completed. Generated", totalSchema, "Schema(s)")
	fmt.Println(schemas)

	var tableNames []string

	var tables []internalTable

	if g.GenerateTables {
		for i := 0; i < g.NumberOfTables; i++ {
			tableName := namesgenerator.GetRandomName(2)
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(g.MaxColumns-g.MinColumns+1) + g.MinColumns // rand.Intn(n) generates a number in [0, n)

			var table internalTable

			table.TableName = tableName

			var fields []string
			for j := 0; j < randomNumber; j++ {
				fieldName := namesgenerator.GetRandomName(2)
				typeVal := randomizeType()

				table.Fields = append(table.Fields, Field{Name: fieldName, Type: typeVal})
				fields = append(fields, fieldName+" "+typeVal.name())
			}

			tables = append(tables, table)

			tableNames = append(tableNames, "CREATE TABLE "+tableName+" ("+strings.Join(fields, ",")+")")
		}
	}

	var insertQueries []string

	fmt.Println("populating tables")

	if g.PopulateTable {
		for _, t := range tables {
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(g.MaxRows-g.MinRows+1) + g.MinRows // rand.Intn(n) generates a number in [0, n)
			insertQueries = append(insertQueries, t.String(randomNumber))
		}
	}

	if g.DryRun {
		for _, s := range schemas {
			fmt.Println(s)
		}

		for _, t := range tableNames {
			fmt.Println(t)
		}

		for _, i := range insertQueries {
			fmt.Println(i)
		}

	} else {

		fmt.Println("creating everything in the database")

		d, err := g.Connection.NewClient()
		if err != nil {
			return err
		}
		for _, s := range schemas {
			if err = d.Exec(s).Error; err != nil {
				log.Println(err)
			}
		}

		for _, t := range tableNames {
			if err = d.Exec(t).Error; err != nil {
				log.Println(err)
			}
		}

		for _, i := range insertQueries {
			if err = d.Exec(i).Error; err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

type TableGeneration struct {
	Connection
	SchemaName     string `json:"schema_name,omitempty" yaml:"schemaName,omitempty"`
	NumberOfTables int    `json:"number_of_tables,omitempty" yaml:"numberOfTables,omitempty"`
	MinColumns     int    `json:"min_columns,omitempty" yaml:"minColumns,omitempty"`
	MaxColumns     int    `json:"max_columns,omitempty" yaml:"maxColumns,omitempty"`
	PopulateTable  bool   `json:"populate_table,omitempty" yaml:"populateTable,omitempty"`
	DataGeneration
}

func (g TableGeneration) GenerateTables() error {

	return nil
}

type DataGeneration struct {
	Connection
	TableName string `json:"table_name,omitempty" yaml:"tableName,omitempty"`
	MinRows   int    `json:"min_rows,omitempty" yaml:"minRows,omitempty"`
	MaxRows   int    `json:"max_rows,omitempty" yaml:"maxRows,omitempty"`
}

func (g DataGeneration) GenerateData() error {
	return nil
}
