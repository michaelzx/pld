package pld_errs

const CommonBizCode = 1000

var (
	Common           = NewBadRequest(10000, "")
	DataNotExist     = NewBadRequest(10001, "不存在")
	DataAlreadyExist = NewBadRequest(10002, "已存在")
	ParamsNotExist   = NewBadRequest(10003, "缺少参数：")
	ParamsErr        = NewBadRequest(10004, "参数错误：")
	QueryNotExist    = NewBadRequest(10003, "缺少参数：")
	QueryErr         = NewBadRequest(10004, "参数错误：")
	ParamsRequired   = NewBadRequest(10005, "缺少必填项：")
)

type IBizErr interface {
}

type BizErr struct {
	Code    int
	Message string
}
type HttpErr struct {
	Status int
	BizErr
}
