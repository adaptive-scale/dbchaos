/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/adaptive-scale/dbchaos/pkg/config"

	"github.com/spf13/cobra"
)

// runScenarioCmd represents the runScenario command
var runScenarioCmd = &cobra.Command{
	Use:   "runScenario",
	Short: "Execute a scenario",
	Long: `Execute a scenario.
A scenario is a collection of tests that can be run in parallel.

Create a file called scenario.yaml with the following content:

dbType: mysql
connection: "root:root@tcp(host:port)/db"
scenarios:
  - query: select * from information_schema.statistics # (for MongoDB, query must be valid JSON ex: '{"insert": "users", "documents": [{ "user": "abc123", "status": "A" }]}')
	parallelRuns: 10000
	runFor: 15m
  - query: |
      SELECT table_schema "Database",
	  ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) "Size (MB)"
	  FROM information_schema.tables
	  GROUP BY table_schema;
	parallelRuns: 10000
	runFor: 15m
dbName: users   #(MongoDB only)

Run as follows:
dbchaos runScenario
`,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := os.ReadFile("./scenario.yaml")
		if err != nil {
			log.Fatal(err)
		}
		s1 := config.ParseScenario(d)
		s1.Start()
	},
}

func init() {
	rootCmd.AddCommand(runScenarioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runScenarioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runScenarioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
