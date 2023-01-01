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
- combination
  - Must
  - MustNot
  - Should
- create
  - AddDoc 添加数据
  - AddDocById 添加数据 指定id
  - CreateIndex 创建一个index, mapping看情况而定.
- delete
  - DeleteDoc 删除文档
  - DelIndex 删除Index
  - DelIndexByID 删除指定id数据
- put
  - Put 上传数据 
- select
  - QueryDoc  Query doc
  - 查询数据.
  - QueryByID 根据id查询文档数据
  - MultiGet 通过多个Id批量查询文档
  - MultiGetByTerms 通过terms查询语法实现，多值查询效果. 通过terms实现SQL的in查询
- update
  - UpdateDoc 条件更新文档
  - Update 更新数据
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