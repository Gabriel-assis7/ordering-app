package main

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Order",
	Short: "Order Service Application",
	Long:  "This is the Order Service Application built with Go.",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: implement application logic
	},
}

func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle(
			"Order Service",
			pterm.BgCyan.ToStyle(),
		),
	).Render()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
