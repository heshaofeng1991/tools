package response

var (
	Succ      = &Codes{Code: 0, DescCh: "ok", DescEn: "ok", Msg: ""}
	Fail      = &Codes{Code: 1, DescCh: "操作失败", DescEn: "operation failed", Msg: ""}
	ParamErr  = &Codes{Code: 2, DescCh: "参数错误", DescEn: "param error", Msg: ""}
	SystemErr = &Codes{Code: 3, DescCh: "系统错误", DescEn: "system error", Msg: ""}
	TokenErr  = &Codes{Code: 401, DescCh: "token失效", DescEn: "key is invalid", Msg: ""}
)
