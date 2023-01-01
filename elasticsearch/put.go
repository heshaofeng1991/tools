package elasticsearch

import (
	"fmt"
)

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
