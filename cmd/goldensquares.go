package cmd

import (
        "github.com/spf13/cobra"
        "fmt"
        "github.com/luphord/algoristic/goldensquares"
)

// goldensquaresCmd represents the version command
var goldensquaresCmd = &cobra.Command{
        Use:   "golden-squares",
        Short: "Create an image of golden squares.",
        Long:  `Golden squares long description ToDo.`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Golden ratio:", goldensquares.GoldenRatio)
        },
}

func init() {
        rootCmd.AddCommand(goldensquaresCmd)
}