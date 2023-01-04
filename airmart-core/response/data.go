package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Data struct {
	Code  int         `json:"code"`
	Desc  string      `json:"desc"`
	Data  interface{} `json:"data"`
	Trace string      `json:"trace_id"` //链路id
}

type Codes struct {
	Code   int
	DescCh string
	DescEn string
	Msg    string
	LogErr error
}

type Empty struct{}

// List 列表格式
type List struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

func ListData(data interface{}, total int64) List {
	return List{List: data, Total: total}
}

func Failed(ctx *gin.Context, data Data, err error) {
	out := data
	output(ctx, out, err)
}

func Success(ctx *gin.Context, data interface{}) {
	out := Data{}
	out.Data = data
	output(ctx, out)
}

func output(ctx *gin.Context, data Data, err ...error) {
	for _, e := range err {
		zap.L().Error("response err", zap.Error(e))
		data.Desc += " err:" + e.Error()
	}
	ctx.JSON(http.StatusOK, data)
}
