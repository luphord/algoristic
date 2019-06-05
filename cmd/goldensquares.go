package cmd

import (
	"fmt"

	"github.com/luphord/algoristic/goldensquares"
	"github.com/spf13/cobra"
)

// goldensquaresCmd represents the version command
var goldensquaresCmd = &cobra.Command{
	Use:   "golden-squares",
	Short: "Create an image of golden squares.",
	Long:  `Golden squares long description ToDo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golden ratio:", goldensquares.GoldenRatio)
		goldensquares.TestSvg()
	},
}

func init() {
	rootCmd.AddCommand(goldensquaresCmd)
}
