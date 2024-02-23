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

type SchemaGeneration struct {
	// Name of the generation
	Connection Connection             `json:"connection,omitempty" yaml:"connection,omitempty"`
	Schema     StaticSchemaGeneration `json:"schema,omitempty" yaml:"schema,omitempty"`
	DryRun     bool                   `json:"dry_run,omitempty" yaml:"dryRun,omitempty"`
	Tables     TableGeneration        `json:"tables,omitempty" yaml:"tables,omitempty"`
	Rows       DataGeneration         `json:"rows,omitempty" yaml:"rows,omitempty"`
}

type StaticSchemaGeneration struct {
	NumberOfSchema int    `json:"number_of_schema,omitempty" yaml:"numberOfSchema,omitempty"`
	GenerateTables bool   `json:"generate_tables,omitempty" yaml:"generateTables,omitempty"`
	Language       string `json:"language,omitempty" yaml:"language,omitempty"`
	SchemaName     string `json:"schema_name,omitempty" yaml:"schemaName,omitempty"`
	DryRun         bool   `json:"dry_run,omitempty" yaml:"dryRun,omitempty"`
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

type SchemaName struct {
	Name string
}

func (s *SchemaName) String() string {
	return s.Name
}

func (s *SchemaName) Create() string {
	return "CREATE SCHEMA " + s.Name
}

func (g SchemaGeneration) GenerateSchema() error {

	totalSchema := g.Schema.NumberOfSchema
	var schemas []SchemaName

	for i := 0; i < totalSchema; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(10)
		name := namesgenerator.GetRandomName(a)
		schemas = append(schemas, SchemaName{Name: name})
	}

	fmt.Println("Schema SchemaGeneration Completed. Generated", totalSchema, "Schema(s)")
	fmt.Println(schemas)

	var tableNames []string

	var tables []internalTable

	if g.Schema.GenerateTables {
		if g.Schema.SchemaName == "" {
			for _, s := range schemas {
				for i := 0; i < g.Tables.NumberOfTables; i++ {

					rand.Seed(time.Now().UnixNano())
					a := rand.Intn(10)
					tableName := namesgenerator.GetRandomName(a)
					rand.Seed(time.Now().UnixNano())
					randomNumber := rand.Intn(g.Tables.MaxColumns-g.Tables.MinColumns+1) + g.Tables.MinColumns // rand.Intn(n) generates a number in [0, n)

					var table internalTable

					table.TableName = s.String() + "." + tableName

					var fields []string
					for j := 0; j < randomNumber; j++ {

						rand.Seed(time.Now().UnixNano())
						a := rand.Intn(10)

						fieldName := namesgenerator.GetRandomName(a)
						typeVal := randomizeType()

						table.Fields = append(table.Fields, Field{Name: fieldName, Type: typeVal})
						fields = append(fields, fieldName+" "+typeVal.name())
					}

					tables = append(tables, table)

					tableNames = append(tableNames, "CREATE TABLE "+s.String()+"."+tableName+" ("+strings.Join(fields, ",")+")")
				}
			}
		} else {
			for i := 0; i < g.Tables.NumberOfTables; i++ {

				rand.Seed(time.Now().UnixNano())
				a := rand.Intn(10)

				tableName := namesgenerator.GetRandomName(a)
				rand.Seed(time.Now().UnixNano())
				randomNumber := rand.Intn(g.Tables.MaxColumns-g.Tables.MinColumns+1) + g.Tables.MinColumns // rand.Intn(n) generates a number in [0, n)

				var table internalTable

				table.TableName = g.Schema.SchemaName + "." + tableName

				var fields []string
				for j := 0; j < randomNumber; j++ {

					rand.Seed(time.Now().UnixNano())
					a := rand.Intn(10)

					fieldName := namesgenerator.GetRandomName(a)
					typeVal := randomizeType()

					table.Fields = append(table.Fields, Field{Name: fieldName, Type: typeVal})
					fields = append(fields, fieldName+" "+typeVal.name())
				}

				tables = append(tables, table)
				tableNames = append(tableNames, "CREATE TABLE "+g.Schema.SchemaName+"."+tableName+" ("+strings.Join(fields, ",")+")")
			}
		}
	}

	var insertQueries []string

	fmt.Println("populating tables")

	if g.Tables.PopulateTable {
		for _, t := range tables {
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(g.Rows.MaxRows-g.Rows.MinRows+1) + g.Rows.MinRows // rand.Intn(n) generates a number in [0, n)
			insertQueries = append(insertQueries, t.String(randomNumber))
		}
	}

	if g.DryRun {
		for _, s := range schemas {
			fmt.Println(s.Create())
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
			if err = d.Exec(s.Create()).Error; err != nil {
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
	Connection     Connection `json:"connection,omitempty" yaml:"connection,omitempty"`
	NumberOfTables int        `json:"number_of_tables,omitempty" yaml:"numberOfTables,omitempty"`
	MinColumns     int        `json:"min_columns,omitempty" yaml:"minColumns,omitempty"`
	MaxColumns     int        `json:"max_columns,omitempty" yaml:"maxColumns,omitempty"`
	PopulateTable  bool       `json:"populate_table,omitempty" yaml:"populateTable,omitempty"`
}

func (g TableGeneration) GenerateTables() error {

	return nil
}

type DataGeneration struct {
	Connection Connection `json:"connection,omitempty" yaml:"connection,omitempty"`
	TableName  string     `json:"table_name,omitempty" yaml:"tableName,omitempty"`
	MinRows    int        `json:"min_rows,omitempty" yaml:"minRows,omitempty"`
	MaxRows    int        `json:"max_rows,omitempty" yaml:"maxRows,omitempty"`
}

func (g DataGeneration) GenerateData() error {
	return nil
}
