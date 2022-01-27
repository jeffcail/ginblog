package types

type Codes struct {
	SUCCESS  uint
	FAILED   uint
	GENERATETOKEN uint
	NOAUTH uint
	AUTHFAILED uint
	AUTHFORMATERROR uint
	INVALIDTOKEN uint
	NOSUCHID uint
	CREATEUSERFAILED uint
	LCAKPARAMETERS uint
	CONVERTFAILED uint
	NOSUCHNAME uint
	EXISTSNAME uint
	Message  map[uint]string
}

var ApiCode = &Codes{
	SUCCESS: 200,
	FAILED: 0,
	AUTHFAILED: 4001,
	GENERATETOKEN: 4002,
	NOAUTH: 4003,
	AUTHFORMATERROR: 4004,
	INVALIDTOKEN: 4005,
	NOSUCHID: 1001,
	CREATEUSERFAILED: 2001,
	LCAKPARAMETERS: 3001,
	CONVERTFAILED: 3002,
	NOSUCHNAME: 5001,
	EXISTSNAME: 5002,
}

func init()  {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS: "成功",
		ApiCode.FAILED:  "失败",
		ApiCode.GENERATETOKEN: "生成Token失败",
		ApiCode.AUTHFAILED: "鉴权失败",
		ApiCode.NOAUTH: "请求头中auth为空",
		ApiCode.AUTHFORMATERROR: "请求头中auth格式有误",
		ApiCode.INVALIDTOKEN:"无效的Token",
		ApiCode.NOSUCHID: "id不存在",
		ApiCode.CREATEUSERFAILED: "用户创建失败",
		ApiCode.LCAKPARAMETERS: "缺少参数",
		ApiCode.CONVERTFAILED: "参数类型转换报错",
		ApiCode.NOSUCHNAME: "根据名称查不到数据",
		ApiCode.EXISTSNAME: "名称重复",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
