package cmd

import (
	"os"

	"github.com/luphord/algoristic/goldensquares"
	"github.com/spf13/cobra"
)

// goldensquaresCmd represents the version command
var goldensquaresCmd = &cobra.Command{
	Use:   "golden-squares",
	Short: "Create an image of golden squares.",
	Long:  `Golden squares long description ToDo.`,
	Run: func(cmd *cobra.Command, args []string) {
		goldensquares.GoldenSquares(os.Stdout, 600)
	},
}

func init() {
	rootCmd.AddCommand(goldensquaresCmd)
}
