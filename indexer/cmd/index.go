/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "index a given directory into a zincsearch server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("index called")
	},
}

func init() {
	indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(indexCmd)

}
