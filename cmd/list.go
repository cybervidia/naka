/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package cmd

import (
	"github.com/cybervidia/naka/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved entries",
	Long:  `Displays a list of all saved entries by their unique names. Passwords are not shown.`,
	Run: func(cmd *cobra.Command, args []string) {

		userTagFlag, err := cmd.Flags().GetString("tag")
		if err != nil {
			panic(err)
		}

		db.ListSecret(userTagFlag)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringP("tag", "t", "", "add a tag to categorize your password/secret")
}
