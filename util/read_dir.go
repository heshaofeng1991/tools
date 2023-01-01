package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 读取目录并返回排好序的文件和子目录名（ []os.FileInfo ）
func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}

			fmt.Println(info.Name(), "\\")

			listAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}

			fmt.Println(info.Name())
		}
	}
}
