package sync

import (
	"github.com/spf13/cobra"
	"gsync/compare"
)

var (
	destPath string
)

func NewCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "check",
		Short:                      "check the localFilesMd5 Values",
		SuggestionsMinimumDistance: 1,
		SuggestFor:                 []string{"sync"},
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
	cmd.Flags().StringVarP(&destPath, "destPath", "d", "", "check /path/to/file")
	cmd.MarkFlagRequired("destPath")
	return cmd
}

func Run() {
	compare.ShowFileMd5List(destPath)
}
