/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/adaptive-scale/dbchaos/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates synthetic data for generating schema, tables and data",
	Long:  `generates synthetic data for generating schema, tables and dat`,
	Run: func(cmd *cobra.Command, args []string) {

		d, err := os.ReadFile("config.yaml")
		if err != nil {
			log.Fatal(err)
		}

		schema, err := cmd.Flags().GetBool("schema")
		if err != nil {
			log.Fatal(err)
		}

		table, err := cmd.Flags().GetBool("table")
		if err != nil {
			log.Fatal(err)
		}

		data, err := cmd.Flags().GetBool("data")
		if err != nil {
			log.Fatal(err)
		}

		if schema {
			var schemaConfig config.StaticSchemaGeneration
			if err := yaml.Unmarshal(d, &schemaConfig); err != nil {
				log.Fatal(err)
			}
			err = schemaConfig.GenerateSchema()
			if err != nil {
				log.Fatal(err)
			}
		} else if table {
			var tableConfig []config.TableGeneration
			if err := yaml.Unmarshal(d, &tableConfig); err != nil {
				log.Fatal(err)
			}

			for _, t := range tableConfig {
				err = t.GenerateTables()
				if err != nil {
					log.Fatal(err)
				}
			}

		} else if data {
			var dataConfig []config.DataGeneration
			if err := yaml.Unmarshal(d, &dataConfig); err != nil {
				log.Fatal(err)
			}

			for _, d := range dataConfig {
				err = d.GenerateData()
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			log.Fatal("Please specify the type of generation")
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	generateCmd.PersistentFlags().Bool("schema", false, "Config yaml file to use")
	generateCmd.PersistentFlags().Bool("table", false, "Config yaml file to use")
	generateCmd.PersistentFlags().Bool("data", false, "Config yaml file to use")
}
