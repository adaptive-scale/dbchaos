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

		var schemaConfig config.SchemaGeneration
		if err := yaml.Unmarshal(d, &schemaConfig); err != nil {
			log.Fatal(err)
		}
		err = schemaConfig.GenerateSchema()
		if err != nil {
			log.Fatal(err)
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

}
