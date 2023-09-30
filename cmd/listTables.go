/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"dynamodbtinker/ddb"
)

// listTablesCmd represents the listTables command
var listTablesCmd = &cobra.Command{
	Use:   "listTables",
	Short: "List the tables in ddb",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listTables called")

		var db ddb.Ddb
		portFlag, err := cmd.Flags().GetInt("port")
		if err != nil {
			fmt.Println(err)
		}

		if portFlag != 0 {
			db = ddb.SetupDDB(portFlag)
		} else {
			db = ddb.SetupDDB(8000)
		}

		tables, err := db.ListTables()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(tables)

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
