package elasticsearch

import (
	"fmt"
	"log"
)

// AddDoc 添加数据.
func (es *Client) AddDoc(index, desc string, data interface{}) (bool, error) {
	// 添加索引.
	_, err := es.CreateIndex(index, desc)
	if err != nil {
		log.Fatal("创建索引失败", err)
	}

	// 添加doc 先获取index再为index添加index.
	res, err := es.Client.Index().
		Index(index).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		return false, err
	}

	fmt.Println("添加数据成功:", res)

	return true, nil
}

// AddDocById 添加数据 指定id.
func (es *Client) AddDocById(id, index, desc string, data interface{}) (bool, error) {
	// 添加索引.
	_, err := es.CreateIndex(index, desc)
	if err != nil {
		log.Fatal("创建索引失败", err)
	}

	// 添加doc 先获取index再为index添加index.
	res, err := es.Client.Index().
		Index(index).
		BodyJson(data).
		Id(id).
		Do(ctx)
	if err != nil {
		return false, err
	}

	fmt.Println("基于ID添加数据成功:", res)

	return true, nil
}

// CreateIndex 创建一个index, mapping看情况而定. 在7.x后强制要求只能有一个类型，就是 _doc  ---> /{index}/_doc/{id}.
func (es *Client) CreateIndex(index, mapping string) (bool, error) {
	// 判断索引是否存在.
	exists, err := es.Client.IndexExists(index).Do(ctx)
	if err != nil {
		fmt.Sprintf("<CreateIndex> some error occurred when check exists, index: %s, err:%s", index, err.Error())

		return false, err
	}

	if !exists {
		createIndex, err := es.Client.CreateIndex(index).Body(mapping).Do(ctx)
		if err != nil {
			fmt.Sprintf("<CreateIndex> some error occurred when create. index: %s, err:%s", index, err.Error())

			return false, err
		}

		if !createIndex.Acknowledged {
			// Not acknowledged.
			fmt.Sprintf("<CreateIndex> Not acknowledged, index: %s", index)

			return false, nil
		}
	}

	return true, nil
}
