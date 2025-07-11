/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "naka",
	Short: "naka banner",
	Long: `naks is a simple and secure command-line password manager written in Go.  
It lets you add, retrieve, update, and delete passwords stored locally in an encrypted SQLite database.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("                      oooo                  \n" +
			"                      `888                  \n" +
			"ooo. .oo.    .oooo.    888  oooo   .oooo.   \n" +
			"`888P\"Y88b  `P  )88b   888 .8P'   `P  )88b  \n" +
			" 888   888   .oP\"888   888888.     .oP\"888  \n" +
			" 888   888  d8(  888   888 `88b.  d8(  888  \n" +
			"o888o o888o `Y888\"\"8o o888o o888o `Y888\"\"8o \n" +
			"                                            \n" +
			"        [ 中 ] naka\n" +
			"     \"The key is inside.\"")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.naka.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
