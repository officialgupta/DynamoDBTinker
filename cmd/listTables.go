/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listTablesCmd represents the listTables command
var listTablesCmd = &cobra.Command{
	Use:   "listTables",
	Short: "List the tables in ddb",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := connect(cmd)
		if err != nil {
			fmt.Println(err)
		}

		tables, err := db.ListTables()
		if err != nil {
			fmt.Println(err)
		}

		if len(*tables) == 0 {
			fmt.Println("No tables found")
			os.Exit(1)
		}

		for _, table := range *tables {
			fmt.Println(table)
		}
	},
}

func init() {
	rootCmd.AddCommand(listTablesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listTablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listTablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
