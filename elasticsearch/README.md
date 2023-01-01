# ElasticSearch

- init
```go
// Connect ElasticSearch连接.
func Connect(host, name, password string) (*Client, error) {
	client, err := elastic.NewClient(
		// host是可变参数 http://10.0.1.1:9200", "http://10.0.1.2:9200.
		elastic.SetURL(host),
		// 基于http base auth验证机制的账号和密码. 这里跟上面的host应该都是全局配置.
		elastic.SetBasicAuth(name, password),
		// 启用gzip压缩.
		elastic.SetGzip(true),
		// 设置监控检查时间间隔.
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置错误日志输出.
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出.
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return nil, err
	}

	info, code, err := client.Ping(host).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Client returned with code %d and version %s\n", code, info.Version.Number)

	esVersion, err := client.ElasticsearchVersion(host)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Client version %s\n", esVersion)

	return &Client{
		Client: client,
	}, nil
}
```
- elasticsearch
```go
func New(conf Elasticsearch) (client *elastic.Client, err error) {
	client, err = elastic.NewSimpleClient(
		elastic.SetURL(conf.Address...),                            // 服务地址
		elastic.SetBasicAuth(conf.Username, conf.Password),         // 账号密码
		elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)), // 设置错误日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),  // 设置info日志输出
	)
	
	return
}
```
- aggregation <br>
    桶聚合 bucket 桶聚合通常用于对数据分组，然后分组内的数据可以使用指标聚合汇总数据。
  - Terms聚合
  - Histogram聚合
  - Date histogram聚合
  - Range聚合
  - 嵌套聚合 
```go
/* 桶聚合 bucket 桶聚合通常用于对数据分组，然后分组内的数据可以使用指标聚合汇总数据。
    1. Terms聚合
    2. Histogram聚合
    3. Date histogram聚合
    4. Range聚合
    5. 嵌套聚合
*/

func (es *Client) Terms(index, aggName, field string) (interface{}, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName, // 聚合条件名称 比如 "shop"
			elastic.NewTermsAggregation().Field(field), // 设置统计字段 "shop_id"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Terms函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Terms(aggName)
	if !isFind {
		return nil, fmt.Errorf("terms not find data")
	}

	return agg.Buckets, nil
}

func (es *Client) Histogram(index, aggName, field string, interval float64) (interface{}, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName, // 聚合条件名称 比如 "price"
			elastic.NewHistogramAggregation().Field(field).Interval(interval), // 设置统计字段 "prices"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Histogram函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Histogram(aggName)
	if !isFind {
		return nil, fmt.Errorf("histogram not find data")
	}

	return agg.Buckets, nil
}

func (es *Client) DateHistogram(index, aggName, intervalName, format, field string) (interface{}, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName,
			elastic.NewDateHistogramAggregation().
				Field(field).                   // "date"
				CalendarInterval(intervalName). // 分组间隔：month代表每月、支持minute（每分钟）、hour（每小时）、day（每天）、week（每周）、year（每年)
				Format(format),                 // "yyyy-MM-dd"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用DateHistogram函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.DateHistogram(aggName)
	if !isFind {
		return nil, fmt.Errorf("DateHistogram not find data")
	}

	return agg.Buckets, nil
}

func (es *Client) Range(index, aggName, field string, from, to float64) (interface{}, error) {
	res, err := es.Client.Search().
		Index(index).
		Query(elastic.NewMatchAllQuery()).
		Aggregation(
			aggName,
			elastic.NewRangeAggregation().
				Field(field).           // 根据price字段分桶
				AddUnboundedFrom(from). // 范围配置, 0 - 100
				AddRange(from, to).     // 范围配置, 100 - 200
				AddUnboundedTo(to),     // 范围配置，> 200的值
		).
		Size(0).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	// 使用Range函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Range(aggName)
	if !isFind {
		return nil, fmt.Errorf("range not find data")
	}

	return agg.Buckets, nil
}

func (es *Client) Aggregation(index, aggName, termsField, sumField string) (interface{}, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName,
			elastic.NewTermsAggregation().Field(termsField).
				SubAggregation(aggName, elastic.NewSumAggregation().Field(sumField)),
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 嵌套聚合
	agg, isFind := res.Aggregations.Terms(aggName)
	if !isFind {
		return nil, fmt.Errorf("ValueCount not find data")
	}

	return agg.Buckets, nil
}

/* 指标聚合 metrics ES指标聚合，就是类似SQL的统计函数，指标聚合可以单独使用，也可以跟桶聚合一起使用
    1. Value Count 值聚合，主要用于统计文档总数，类似SQL的count函数.
    2. Cardinality
        基数聚合，也是用于统计文档的总数，跟Value Count的区别是，基数聚合会去重，不会统计重复的值，类似SQL的count(DISTINCT 字段)用法
        基数聚合是一种近似算法，统计的结果会有一定误差，不过性能很好
    3. Avg 求平均值
    4. Sum 求和计算
    5. Max 求最大值
    6. Min Min
*/

func (es *Client) ValueCount(index, aggName, field string) (*float64, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName, // 聚合条件名称 比如 "total"
			elastic.NewValueCountAggregation().Field(field), // 设置统计字段 "order_id"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用ValueCount函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.ValueCount(aggName)
	if !isFind {
		return nil, fmt.Errorf("ValueCount not find data")
	}

	return agg.Value, nil
}

func (es *Client) Cardinality(index, aggName, field string) (*float64, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName, // 聚合条件名称 比如 "total"
			elastic.NewCardinalityAggregation().Field(field), // 设置统计字段 "order_id"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Cardinality函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Cardinality(aggName)
	if !isFind {
		return nil, fmt.Errorf("cardinality not find data")
	}

	return agg.Value, nil
}

func (es *Client) Avg(index, aggName, field string) (*float64, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName,                                  // 聚合条件名称 比如 "average_price"
			elastic.NewAvgAggregation().Field(field), // 设置统计字段 "price"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Avg函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Avg(aggName)
	if !isFind {
		return nil, fmt.Errorf("cardinality not find data")
	}

	return agg.Value, nil
}

func (es *Client) Sum(index, aggName, field string) (*float64, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName,                                  // 聚合条件名称 比如 "total_price"
			elastic.NewSumAggregation().Field(field), // 设置统计字段 "price"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Sum函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Sum(aggName)
	if !isFind {
		return nil, fmt.Errorf("cardinality not find data")
	}

	return agg.Value, nil
}

func (es *Client) Max(index, aggName, field string) (*float64, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName,                                  // 聚合条件名称 比如 "max_price"
			elastic.NewMaxAggregation().Field(field), // 设置统计字段 "price"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Max函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Max(aggName)
	if !isFind {
		return nil, fmt.Errorf("cardinality not find data")
	}

	return agg.Value, nil
}

func (es *Client) Min(index, aggName, field string) (*float64, error) {
	res, err := es.Client.Search().
		Index(index).                      // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation(
			aggName,                                  // 聚合条件名称 比如 "min_price"
			elastic.NewMinAggregation().Field(field), // 设置统计字段 "price"
		).       // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)  // 执行请求
	if err != nil {
		return nil, err
	}

	// 使用Min函数和前面定义的聚合条件名称，查询结果
	agg, isFind := res.Aggregations.Min(aggName)
	if !isFind {
		return nil, fmt.Errorf("cardinality not find data")
	}

	return agg.Value, nil
}
```
- combination
  - Must
  - MustNot
  - Should
```go
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
```
- create
  - AddDoc 添加数据
  - AddDocById 添加数据 指定id
  - CreateIndex 创建一个index, mapping看情况而定.
```go
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
```
- delete
  - DeleteDoc 删除文档
  - DelIndex 删除Index
  - DelIndexByID 删除指定id数据
```go
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
```
- put
  - Put 上传数据 
```go
// Put 上传数据.
func (es *Client) Put(index string, bodyJSON interface{}) bool {
	_, err := es.Client.Index().
		Index(index).
		BodyJson(bodyJSON).
		Do(ctx)
	if err != nil {
		// Handle error.
		fmt.Sprintf("<Put> some error occurred when put.  err:%s", err.Error())

		return false
	}

	return true
}
```
- select
  - QueryDoc  Query doc
  - 查询数据.
  - QueryByID 根据id查询文档数据
  - MultiGet 通过多个Id批量查询文档
  - MultiGetByTerms 通过terms查询语法实现，多值查询效果. 通过terms实现SQL的in查询
```go
/*
	index:"index"      索引
	field:"name,age"   要查询的字段
	filter:*TermQuery  查询规则
	sort:"age asc"     排序规则
	page:0             页数
	limit:10           条目数量
*/

// QueryDoc  Query doc.
func (es *Client) QueryDoc(index string, field string, filter elastic.Query, sort string, page int, limit int) (interface{}, error) {
	res, err := es.query(index, field, filter, sort, page, limit)

	strD, _ := json.Marshal(res)

	if err != nil {
		fmt.Println("失败:", err)
		return nil, nil
	}

	fmt.Println("查询数据成功")

	return string(strD), nil
}

// 查询数据.
func (es *Client) query(index string, field string, filter elastic.Query, sort string, page, limit int) (*elastic.SearchResult, error) {
	// 分页数据处理.
	isAsc := true

	if sort != "" {
		sortSlice := strings.Split(sort, " ")
		sort = sortSlice[0]
		if sortSlice[1] == "desc" {
			isAsc = false
		}
	}

	// 查询位置处理.
	if page <= 1 {
		page = 1
	}

	// 返回字段处理.
	fsc := elastic.NewFetchSourceContext(true)

	if field != "" {
		fieldSlice := strings.Split(field, ",")
		if len(fieldSlice) > 0 {
			for _, v := range fieldSlice {
				fsc.Include(v)
			}
		}
	}

	// 开始查询位置
	fromStart := (page - 1) * limit

	res, err := es.Client.Search().
		Index(index).            // 设置索引名.
		FetchSourceContext(fsc). // 指示应如何获取 _source.
		Query(filter).           // 设置查询条件， 这里也可以用elastic.NewTermQuery("Author", "tizi")代替，创建term查询条件，用于精确查询.
		Sort(sort, isAsc).       // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序.
		From(fromStart).         // 设置分页参数 - 起始偏移量，从第0行记录开始.
		Size(limit).             // 设置分页参数 - 每页大小.
		Pretty(true).            // 查询结果返回可读性较好的JSON格式.
		Do(ctx)                  // 执行请求.
	if err != nil {
		return nil, err
	}

	//if res.TotalHits() > 0 {
	//	// 查询结果不为空，则遍历结果.
	//	var b1 struct{}
	//	// 通过Each方法，将es结果的json结构转换成struct对象.
	//	for _, item := range res.Each(reflect.TypeOf(b1)) {
	//		// 转换成结构体对象.
	//		if t, ok := item.(struct{}); ok {
	//			fmt.Println(t)
	//		}
	//	}
	//}

	return res, nil

}

// QueryByID 根据id查询文档数据.
func (es *Client) QueryByID(index, id string) (interface{}, error) {
	res, err := es.Client.Get().
		Index("weibo"). // 指定索引名.
		Id("1").        // 设置文档id.
		Do(ctx)         // 执行请求.
	if err != nil {
		// Handle error.
		return nil, err
	}

	if res.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", res.Id, res.Version, res.Index)
	}

	//// 手动将文档内容转换成go struct对象.
	//msg := struct{}{}
	//// 提取文档内容，原始类型是json数据.
	//data, _ := res.Source.MarshalJSON()
	//// 将json转成struct结果.
	//_ = json.Unmarshal(data, &msg)
	//// 打印结果
	//fmt.Println(msg)

	return res, nil
}

// MultiGet 通过多个Id批量查询文档.
func (es *Client) MultiGet(index string, id ...string) (interface{}, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("empty query id")
	}

	multiGet := es.Client.MultiGet()
	for _, val := range id {
		multiGet.Add(elastic.NewMultiGetItem(). // 通过NewMultiGetItem配置查询条件.
							Index(index). // 设置索引名.
							Id(val))      // 设置文档id.
	}

	result, err := multiGet.Do(ctx)
	if err != nil {
		fmt.Sprintf("MultiGet err %v", err.Error())

		return nil, err
	}

	//// 遍历文档.
	//for _, doc := range result.Docs {
	//	// 转换成struct对象.
	//	var content struct{}
	//	tmp, _ := doc.Source.MarshalJSON()
	//	if err := json.Unmarshal(tmp, &content); err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(content)
	//}

	return result, nil
}

// MultiGetByTerms 通过terms查询语法实现，多值查询效果. 通过terms实现SQL的in查询
func (es *Client) MultiGetByTerms(name string, page, limit int, values ...string) (interface{}, error) {
	// 查询位置处理.
	if page <= 1 {
		page = 1
	}

	// 开始查询位置
	start := (page - 1) * limit

	res, err := es.Client.Search().
		Index("blogs").                             // 设置索引名.
		Query(elastic.NewTermsQuery(name, values)). // 设置查询条件.
		Sort("Created", true).                      // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序.
		From(start).                                // 设置分页参数 - 起始偏移量，从第0行记录开始.
		Size(10).                                   // 设置分页参数 - 每页大小.
		Do(ctx)                                     // 执行请求.
	if err != nil {
		return nil, err
	}

	return res, nil
}
```
- update
  - UpdateDoc 条件更新文档
  - Update 更新数据
```go
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
```
- other
  - GetSourceByID 基于ID查询数据
  - GetTaskMsg 获取任务消息
  - GetTaskLogCount 获取任务数
```go
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
```
### Elastic操作
#### Index
#### Get
#### Delete
#### Delete By Query
#### Update
#### Update By Query
#### Multi Get
#### Bulk
#### Reindex
#### Term Vectors
#### Multi Term Vectors
#### ?Refresh
#### Optimistic Concurrency Control