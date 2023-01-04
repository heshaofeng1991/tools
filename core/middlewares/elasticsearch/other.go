package elasticsearch

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

// GetSourceByID 基于ID查询数据.
func (es *Client) GetSourceByID(index, eSid string) (*elastic.GetResult, error) {
	source, err := es.Client.Get().
		Index(index).
		Id(eSid).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return source, nil
}

// GetTaskMsg 获取任务消息.
func (es *Client) GetTaskMsg(index, name, startTime, endTime, keyword string, size, page int) (*elastic.SearchResult, error) {
	boolQ := elastic.NewBoolQuery()

	if keyword == "" {
		boolQ.Filter(elastic.NewRangeQuery(name).Gte(startTime), elastic.NewRangeQuery(name).Lte(endTime))

		res, err := es.Client.Search(index).
			Query(boolQ).
			Size(size).
			From((page - 1) * size).
			Do(ctx)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	keys := fmt.Sprintf("name:*%s*", keyword)
	boolQ.Filter(elastic.NewRangeQuery(name).Gte(startTime), elastic.NewRangeQuery(name).Lte(endTime), elastic.NewQueryStringQuery(keys))

	res, err := es.Client.Search(index).
		Query(boolQ).
		Size(size).
		From((page - 1) * size).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTaskLogCount 获取任务数.
func (es *Client) GetTaskLogCount(index, name, startTime, endTime string) (int, error) {
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewRangeQuery(name).Gte(startTime), elastic.NewRangeQuery(name).Lte(endTime))

	//统计count
	count, err := es.Client.Count(index).Query(boolQ).Do(ctx)
	if err != nil {
		return 0, nil
	}

	return int(count), nil
}
