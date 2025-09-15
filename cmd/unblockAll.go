/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/noktilumo/unblock/roblox"
	"github.com/spf13/cobra"
)

// unblockAllCmd represents the unblockAll command
var unblockAllCmd = &cobra.Command{
	Use:   "unblockAll",
	Short: "Automatically unblock all the blocked users",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		roblox.UnblockAllBlockedUsers()
	},
}

func init() {
	rootCmd.AddCommand(unblockAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unblockAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unblockAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
