/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var goodbyes = []string{
	"Goodbye! (English)",
	"¡Adiós! (Spanish)",
	"Au revoir! (French)",
	"Arrivederci! (Italian)",
	"さようなら! (Japanese)",
	"До свидания! (Russian)",
	"Auf Wiedersehen! (German)",
	"Adeus! (Portuguese)",
	"再见! (Chinese)",
	"안녕히 가세요! (Korean)",
}

// goodbyeCmd represents the goodbye command
var goodbyeCmd = &cobra.Command{
	Use:   "goodbye",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(goodbyes[rand.Intn(len(goodbyes))])
	},
}

func init() {
	greetCmd.AddCommand(goodbyeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goodbyeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goodbyeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
