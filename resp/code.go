package resp

const (
	// 成功
	success = 200
	// 数据列表
	datalist    = 201
	excel       = 202
	multitenant = 203
	fail        = 500
	//	// 异常告警
	warn         = 501
	refuse       = 401
	unauthorized = 401
	//	// 登录失败
	//	authFail = 402
	// 数据校验异常
	badrequest = 400
	validate   = 400
)

var _codes = map[string]int{
	"success":      success,
	"datalist":     datalist,
	"excel":        excel,
	"multitenant":  multitenant,
	"fail":         fail,
	"error":        fail,
	"warn":         warn,
	"refuse":       refuse,
	"unauthorized": unauthorized,
	"badrequest":   badrequest,
	"valid":        validate,
}

func SetCode(name string, c int) {
	_codes[name] = c
}

func Code(name string) int {
	return _codes[name]
}
