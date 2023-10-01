/*
Copyright © 2023 MAYANK GUPTA
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createTableCmd represents the createTable command
var createTableCmd = &cobra.Command{
	Use:   "createTable [name of table] [key of table]",
	Short: "Create a new DynamoDB table",
	Long: `This command will create a new DynamoDB table. 
	It requires a table name to be provided and a single
	string based primary key. 
	This will be used as a partition key for the new table.	
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := connect(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tableName := args[0]
		key := args[1]

		if tableName == "" || key == "" {
			fmt.Println("must provide a name and a key when creating a table")
			os.Exit(1)
		}

		err = db.CreateTable(tableName, key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Created table %v with key %v\n", tableName, key)
	},
}

func init() {
	rootCmd.AddCommand(createTableCmd)

	// Cobra Persistent Flags

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
