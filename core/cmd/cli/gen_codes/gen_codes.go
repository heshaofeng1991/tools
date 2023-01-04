package gen_codes

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

type codes struct {
	Code  int
	key   string
	index int
}

var regFileName = regexp.MustCompile(`^[\d]+`)

func saveJson(filename string, data [][]interface{}) error {
	out := map[string]interface{}{}
	for _, v := range data {
		m := map[string]interface{}{}
		if len(v) == 1 {

		}
		code, err := strconv.Atoi(v[1].(string))
		if err != nil {
			return errors.Wrapf(err, "code:%v", v)
		}
		m["code"] = code
		m["desc_ch"] = v[2]
		if len(v) == 4 {
			m["desc_en"] = v[3]
		}
		out[v[0].(string)] = m
	}

	indent, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return errors.Wrap(err, "解析json失败")
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(indent)
	return err
}

func saveJs(filename string, data [][]interface{}) error {
	if filename == "" {
		return errors.New("文件名为空")
	}
	var body = "export const errCodes = {\r\n"
	for _, v := range data {
		newSlice := make([]interface{}, 0)
		for _, vv := range v {
			newSlice = append(newSlice, strings.Replace(vv.(string), "\n", "", -1))
		}
		body += fmt.Sprintf("  %s: {\r\n    code: %s,\r\n    desc_ch: '%s',\r\n    desc_en: '%s',\n  },\n", newSlice...)
	}
	body += "};"

	f, err := os.Create("./" + filename + ".ts")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(body)
	return err
}

func saveGo(filename string, data [][]interface{}) error {
	if filename == "" {
		return errors.New("文件名为空")
	}

	body := "package response\n\r"
	body += `
import "airmart-core/response"

`
	body += "\n\r"
	body += "var ("

	for _, v := range data {
		body += "\r\n"
		body += fmt.Sprintf(`  %s  =  &response.Codes{Code: %s, DescCh: "%s", DescEn: "%s"} `, v...)
	}

	body += "\n)"

	f, err := os.Create(filename + ".go")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(body)
	return err
}

func ExportCodes(file string, outPath string) error {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return errors.Wrap(err, "打开文件失败")
	}
	defer f.Close()

	_, err = os.Stat(outPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(outPath, 0666)
		if err != nil {
			return err
		}
	}
	allArgs := make([][]interface{}, 0)
	list := f.GetSheetList()
	for _, v := range list {
		rows, err := f.GetRows(v)
		if err != nil {
			return errors.Wrapf(err, "读取sheet:%s失败", v)
		}
		if len(rows) <= 5 {
			continue
		}
		A2 := rows[1]
		keys := make([]codes, 0)
		filename := outPath + "/" + regFileName.FindString(v)
		field := rows[4][1:]
		rows = rows[5:]
		for i, row := range rows {
			code, err := strconv.Atoi(row[2])
			if err != nil {
				return errors.Wrap(err, "code转int失败")
			}
			keys = append(keys, codes{Code: code, key: row[1], index: i})
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].Code < keys[j].Code
		})
		rowArgs := make([][]interface{}, 0)
		for _, code := range keys {
			row := rows[code.index][1:]
			args := make([]interface{}, 0)
			for fi, _ := range field {
				if fi == 0 && row[fi] == "" {
					break
				}
				rowValue := ""
				if len(row) > fi {
					rowValue = row[fi]
				}
				args = append(args, rowValue)
			}
			if len(args) != 0 {
				rowArgs = append(rowArgs, args)
			}
		}
		allArgs = append(allArgs, rowArgs...)
		if len(A2) == 0 {
			if err := saveJs(filename, rowArgs); err != nil {
				return errors.Wrapf(err, "保存js文件失败 %s", v)
			}
		} else {
			if err := saveGo(filename, rowArgs); err != nil {
				return errors.Wrapf(err, "保存go文件失败 %s", v)
			}
		}
	}
	return saveJson(outPath+"/err_codes.json", allArgs)
}
