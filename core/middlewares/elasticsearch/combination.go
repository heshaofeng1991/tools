package elasticsearch

import "github.com/olivere/elastic/v7"

func (es *Client) Must(termName, termValue, matchName, matchValue, sort string, page, limit int) (interface{}, error) {
	// 查询位置处理.
	if page <= 1 {
		page = 1
	}

	// 开始查询位置
	start := (page - 1) * limit

	return es.Client.Search().
		Index("blogs"). // 设置索引名
		Query(
			elastic.NewBoolQuery().Must().Must(
				elastic.NewTermQuery(termName, termValue),
				elastic.NewMatchQuery(matchName, matchName),
					)). // 设置查询条件
		Sort(sort, true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(start).      // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(page).       // 设置分页参数 - 每页大小
		Do(ctx)           // 执行请求
}

func (es *Client) MustNot(termName, termValue, sort string, page, limit int) (interface{}, error) {
	// 查询位置处理.
	if page <= 1 {
		page = 1
	}

	// 开始查询位置
	start := (page - 1) * limit

	return es.Client.Search().
		Index("blogs"). // 设置索引名
		Query(
			elastic.NewBoolQuery().Must().MustNot(
				elastic.NewTermQuery(termName, termValue),
					)). // 设置查询条件
		Sort(sort, true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(start).      // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(page).       // 设置分页参数 - 每页大小
		Do(ctx)           // 执行请求
}

func (es *Client) Should(termName, termValue, matchName, matchValue, sort string, page, limit int) (interface{}, error) {
	// 查询位置处理.
	if page <= 1 {
		page = 1
	}

	// 开始查询位置
	start := (page - 1) * limit

	return es.Client.Search().
		Index("blogs"). // 设置索引名
		Query(
			elastic.NewBoolQuery().Must().Should(
				elastic.NewTermQuery(termName, termValue),
				elastic.NewMatchQuery(matchName, matchName),
					)). // 设置查询条件
		Sort(sort, true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(start).      // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(page).       // 设置分页参数 - 每页大小
		Do(ctx)           // 执行请求
}
