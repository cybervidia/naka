/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cybervidia/naka/db"
	"github.com/cybervidia/naka/model"
	"github.com/cybervidia/naka/vault"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		//per il momento, punto ad mvp e inserisco i dati da cmd
		//con il formato:
		//naka add <unique-name> <mail/username> <pwd> <notes>

		if len(args) != 4 {
			fmt.Println("syntax error, use:\nnaka add <unique-name> <mail/username> <pwd> <notes>")
			return
		}
		scrt := model.SecretEntry{
			Name:     args[0],
			Mail:     args[1],
			Password: args[2],
			Note:     args[3],
			// IV:       "stringapercryptare", //non ancora inplementato
		}

		vault.Lock(&scrt)
		db.AddSecret(&scrt)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
