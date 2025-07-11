/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package cmd

import (
	"fmt"

	"github.com/cybervidia/naka/db"
	"github.com/cybervidia/naka/model"
	"github.com/cybervidia/naka/vault"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing password entry",
	Long:  `Updates an existing entry in the database with a new user/email, password, or note. The entry must exist and is identified by its unique name.`,
	Run: func(cmd *cobra.Command, args []string) {
		//per il momento, punto ad mvp e inserisco i dati da cmd
		//con il formato:
		//naka add <unique-name> <mail/username> <pwd> <notes>

		// WARNING forse dovrei chiedere la vecchia password??

		if len(args) != 4 {
			fmt.Println("syntax error, use:\nnaka add <unique-name> <mail/username> <pwd> <notes>")
			return
		}
		scrt := model.SecretEntry{
			Name:     args[0],
			Mail:     args[1],
			Password: args[2],
			Note:     args[3],
		}
		vault.Lock(&scrt)
		db.UpdateSecret(&scrt)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
