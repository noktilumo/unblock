/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/noktilumo/unblock/roblox"
	"github.com/spf13/cobra"
)

// getBlockedCmd represents the getBlocked command
var getBlockedCmd = &cobra.Command{
	Use:   "getBlocked",
	Short: "Print the id's of all the blocked users",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(roblox.FetchAllBlockedUserIds())
	},
}

func init() {
	rootCmd.AddCommand(getBlockedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBlockedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBlockedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
