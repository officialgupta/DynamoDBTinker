/*
Copyright © 2023 MAYANK GUPTA
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteTableCmd represents the deleteTable command
var deleteTableCmd = &cobra.Command{
	Use:   "deleteTable [table to delete]",
	Short: "Delete a DynamoDB table",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := connect(cmd)
		if err != nil {
			fmt.Println(err)
		}

		tableName := args[0]

		if tableName == "" {
			fmt.Println("must provide a name when deleting a table")
			os.Exit(1)
		}

		err = db.DeleteTable(tableName)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("table %v deleted", tableName)
	},
}

func init() {
	rootCmd.AddCommand(deleteTableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
