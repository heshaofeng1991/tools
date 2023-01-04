package elasticsearch

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

// DeleteDoc 删除文档。
func (es *Client) DeleteDoc(index string, filter elastic.Query) (bool, error) {
	res, err := es.Client.DeleteByQuery().
		Query(filter).
		Index(index).
		Do(ctx)
	if err != nil {
		return false, err
	}

	fmt.Println("删除信息成功：", res)

	return true, nil
}

// DelIndex 删除Index.
func (es *Client) DelIndex(index string) bool {
	// Delete an index.
	deleteIndex, err := es.Client.DeleteIndex(index).Do(ctx)
	if err != nil {
		// Handle error.
		fmt.Sprintf("<DelIndex> some error occurred when delete. index: %s, err:%s", index, err.Error())

		return false
	}

	if !deleteIndex.Acknowledged {
		// Not acknowledged.
		fmt.Sprintf("<DelIndex> acknowledged. index: %s", index)

		return false
	}

	return true
}

// DelIndexByID 删除指定id数据.
func (es *Client) DelIndexByID(index, id string) bool {
	del, err := es.Client.Delete().
		Index(index).
		Id(id).
		Do(ctx)
	if err != nil {
		// Handle error.
		fmt.Sprintf("<Del> some error occurred when put.  err:%s", err.Error())

		return false
	}

	fmt.Sprintf("<Del> success, id: %s to index: %s, type %s\n", del.Id, del.Index, del.Type)

	return true
}
