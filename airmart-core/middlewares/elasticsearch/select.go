package elasticsearch

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

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
