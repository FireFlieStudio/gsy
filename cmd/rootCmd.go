package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gsync/cmd/push"
	"gsync/cmd/sync"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gsy",
	Short: "gsy is a sync application",
	Long:  "gsy is a sync application based on TencentCos golang sdk,it's highly convenience",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(sync.NewCheck())
	rootCmd.AddCommand(push.NewPush())
}
