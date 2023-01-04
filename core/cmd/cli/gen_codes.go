package cli

import (
	"core/cmd/cli/gen_codes"
	"log"

	"github.com/spf13/cobra"
)

var (
	codesFile string
	outPath   string
)

// 生成响应code
var genCodesCmd = &cobra.Command{
	Use:   "codes",
	Short: "生成code码",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := gen_codes.ExportCodes(codesFile, outPath)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("导出成功")
		}
	},
}

func init() {
	genCodesCmd.PersistentFlags().StringVar(&codesFile, "file", "./err_codes.xlsx", "文件路径")
	genCodesCmd.PersistentFlags().StringVar(&outPath, "outPath", "./out", "输出文件路径")
}
