/*
Copyright © 2023 MAYANK GUPTA
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete all tables in DynamoDB",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := connect(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tables, err := db.ListTables()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prefix, err := cmd.Flags().GetString("prefix")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, table := range *tables {
			if prefix != "" && !strings.HasPrefix(table, prefix) {
				fmt.Printf("Skip deleting table %v\n", table)
				continue
			}

			err := db.DeleteTable(table)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Table %v deleted\n", table)
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	resetCmd.PersistentFlags().String("prefix", "", "Only delete tables with the given prefix")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
