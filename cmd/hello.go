/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var hellos = []string{
	"Hello! (English)",
	"¡Hola! (Spanish)",
	"Bonjour! (French)",
	"Ciao! (Italian)",
	"こんにちは! (Japanese)",
	"Привет! (Russian)",
	"Guten Tag! (German)",
	"Olá! (Portuguese)",
	"你好! (Chinese)",
	"안녕하세요! (Korean)",
}

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(hellos[rand.Intn(len(hellos))])
	},
}

func init() {
	greetCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
