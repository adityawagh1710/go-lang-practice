package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My awesome CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Cobra!")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greet someone",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

func main() {
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().StringP("name", "n", "", "Name to greet")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
