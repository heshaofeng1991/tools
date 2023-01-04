package elasticsearch

import (
	"errors"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// UpdateDoc 条件更新文档.
func (es *Client) UpdateDoc(index string, filter elastic.Query, data map[string]interface{}) (bool, error) {
	// 修改数据组装.
	if len(data) < 0 {
		return false, errors.New("修改参数不正确")
	}

	scriptStr := ""

	for k := range data {
		scriptStr += "ctx._source." + k + " = params." + k + ";"
	}

	script := elastic.NewScript(scriptStr).Params(data)

	res, err := es.Client.UpdateByQuery(index).
		Query(filter). // 想要更多复杂查询请参照这个 https://www.tizi365.com/archives/858.html.
		Script(script).
		Do(ctx)
	if err != nil {
		return false, err
	}

	fmt.Println("更新数据成功:", res)

	return true, nil
}

// Update 更新数据.
func (es *Client) Update(index, id string, updateMap map[string]interface{}) bool {
	res, err := es.Client.Update().
		Index(index).
		Id(id).
		Doc(updateMap).
		FetchSource(true).
		Do(ctx)
	if err != nil {
		_ = fmt.Sprintf("<Update> some error occurred when update. index:%s, id:%s err:%s", index, id, err.Error())

		return false
	}

	if res == nil {
		fmt.Sprintf("<Update> expected response != nil. index:%s, id:%s", index, id)

		return false
	}

	if res.GetResult == nil {
		fmt.Sprintf("<Update> expected GetResult != nil. index:%s, id:%s", index, id)

		return false
	}

	return true
}
