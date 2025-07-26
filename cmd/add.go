/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package cmd

import (
	"fmt"

	"github.com/cybervidia/naka/db"
	"github.com/cybervidia/naka/model"
	"github.com/cybervidia/naka/vault"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [<unique-name> <mail> <pwd> <notes>]",
	Short: "Add a new password entry",
	Long: `Adds a new password entry to the local encrypted database. The entry must have a unique name, a user or email, the password to store, and an optional note or suggestion.
	Syntax:
	naka add <unique-name> <mail> <pwd> <notes>
	Ex:
	naka add maks info@me.it S3crEt@pWd the_pwd_used_in_dev`,
	Run: func(cmd *cobra.Command, args []string) {

		//naka add <unique-name> <mail> <pwd> <notes>

		//se c'è il flag -p

		//naka add -p <unique-name> <mail> <notes>
		scrt := model.SecretEntry{}
		pwdFlag, err := cmd.Flags().GetBool("pwd")
		if err != nil {
			panic(err)
		}

		if pwdFlag {

			pInput := pterm.DefaultInteractiveTextInput.WithMask("中")
			pFromUsr, err := pInput.Show("中put your secret to seal here")
			if err != nil {
				panic(err)
			}

			scrt = model.SecretEntry{
				Name:     args[0],
				Mail:     args[1],
				Password: pFromUsr,
				Note:     args[2],
			}

		} else {

			//add a warning x hisotry
			pterm.Warning.Println("passwords/secrets you want to store might remain in the shell history.\nI recommend using the -p flag.\nExample:naka add -p <unique_name> <mail> <notes>")

			if len(args) != 4 {
				fmt.Println("Syntax Error, use:\nnaka add <unique-name> <mail/username> <pwd> <notes>")
				return
			}
			scrt = model.SecretEntry{
				Name:     args[0],
				Mail:     args[1],
				Password: args[2],
				Note:     args[3],
			}
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
	addCmd.Flags().BoolP("pwd", "p", false, "ask for password in a secure unseen way\nSyntax:\nnaka add -p <unique-name> <mail> <notes>\nEx:\nnaka add -p maks info@me.it the_pwd_used_in_dev")
}
