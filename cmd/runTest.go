/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/adaptive-scale/dbchaos/pkg/config"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// runTestCmd represents the runTest command
var runTestCmd = &cobra.Command{
	Use:   "runTest",
	Short: "Execute a test on DB",
	Long: `Execute a test on DB.

Create a file name `config.yaml` with the following content:

dbType: postgres
connection: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
query: |
	SELECT pg_database.datname as "Database", pg_size_pretty(pg_database_size(pg_database.datname)) as "Size"
	FROM pg_database;
parallelRuns: 100
runFor: 30m

To run the above config file:

dbchaos runTest config.yaml

`,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := os.ReadFile("./config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		s1 := config.ParseConfiguration(d)
		err = s1.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
