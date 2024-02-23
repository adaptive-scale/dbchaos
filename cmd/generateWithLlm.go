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

// generateWithLlmCmd represents the generateWithLlm command
var generateWithLlmCmd = &cobra.Command{
	Use:   "generateWithLlm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		configFile, _ := cmd.Flags().GetString("config")

		d, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}

		var schemaConfig config.SchemaGenerationWithLLM
		if err := yaml.Unmarshal(d, &schemaConfig); err != nil {
			log.Fatal("invalid schema configuration -", err)
		}

		apiKey := os.Getenv("OPENAI_API_KEY")

		if apiKey == "" {
			log.Fatal("Please set OPENAI_API_KEY environment variable")
		}

		err = schemaConfig.GenerateSchema(apiKey)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateWithLlmCmd)

	generateWithLlmCmd.Flags().String("config", "config.yaml", "config file (default is config.yaml)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateWithLlmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateWithLlmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
