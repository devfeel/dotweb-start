package contract

type ResponseInfo struct {
	RetCode int
	RetMsg  string
	Message interface{}
}

func NewResonseInfo() *ResponseInfo {
	return &ResponseInfo{}
}

func CreateResponse(retCode int, retMsg string, message interface{}) *ResponseInfo {
	return &ResponseInfo{
		RetCode: retCode,
		RetMsg:  retMsg,
		Message: message,
	}
}
