package elasticsearch

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

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
