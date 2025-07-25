/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package cmd

import (
	"github.com/cybervidia/naka/db"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a stored password",
	Long:  `Retrieves and decrypts a password by its unique name and it will be copied to the clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		db.GetSecret(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
