package cmd

import (
	"github.com/officialgupta/DynamoDBTinker/ddb"

	"github.com/spf13/cobra"
)

func connect(cmd *cobra.Command) (*ddb.Ddb, error) {
	var db ddb.Ddb
	portFlag, err := cmd.Flags().GetInt("port")
	if err != nil {
		return nil, err
	}

	if portFlag != 0 {
		db = ddb.SetupDDB(portFlag)
	} else {
		db = ddb.SetupDDB(8000)
	}

	return &db, nil
}
