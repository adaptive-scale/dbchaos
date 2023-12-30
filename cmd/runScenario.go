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

// runScenarioCmd represents the runScenario command
var runScenarioCmd = &cobra.Command{
	Use:   "runScenario",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
