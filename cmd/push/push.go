package push

import (
	"github.com/spf13/cobra"
	"gsync/TencentCos"
)

var (
	bucketName, remoteDstDir, localSrc string
)

func NewPush() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "push",
		Short:                      "push the file to TencentCos bucket",
		SuggestionsMinimumDistance: 1,
		SuggestFor:                 []string{"sync"},
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
	cmd.Flags().StringVarP(&bucketName, "bucketName", "b", "", "push -b bucketName -r /path/to/file -l /remote/path/to/file")
	cmd.Flags().StringVarP(&remoteDstDir, "remoteDstDir", "r", "", "push -b bucketName -r /path/to/file -l /remote/path/to/file")
	cmd.Flags().StringVarP(&localSrc, "localSrc", "l", "", "push -b bucketName -r /path/to/file -l /remote/path/to/file")
	cmd.MarkFlagRequired("bucketName")
	cmd.MarkFlagRequired("remoteDstDir")
	cmd.MarkFlagRequired("localSrc")
	return cmd
}

func Run() {
	fileManager := TencentCos.NewFileManager(bucketName)
	fileManager.DefaultUpload(remoteDstDir, localSrc)
}
