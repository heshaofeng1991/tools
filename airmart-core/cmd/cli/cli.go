package cli

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: `工具命令程序`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Printf("fail on RootCmd.Execute err:%s", err.Error())
	}
}

func init() {
	RootCmd.AddCommand(genCodesCmd) //生成code码
}
