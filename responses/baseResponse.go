package responses

import "go-simple-blog/contracts/statusCodes"

type BaseResponse struct {
	status     bool
	statusCode statusCodes.StatusCode
	message    string
	data       any
}

func (r BaseResponse) GetStatus() bool {
	return r.status
}

func (r BaseResponse) GetStatusCode() statusCodes.StatusCode {
	return r.statusCode
}

func (r BaseResponse) GetMessage() string {
	return r.message
}

func (r BaseResponse) GetData() any {
	return r.data
}

func (r *BaseResponse) SetMessage(msg string) {
	r.message = msg
}

func (r BaseResponse) ErrorIs(code statusCodes.StatusCode) bool {
	if r.statusCode == code {
		return true
	}
	return false
}

func (r BaseResponse) IsFailed() bool {
	if r.status == true {
		return false
	}
	return true
}

func (r BaseResponse) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status":     r.status,
		"statusCode": r.statusCode,
		"message":    r.message,
		"data":       r.data,
	}
}

func New[Res PostResponse | UserResponse](res Res, status bool, statusCode statusCodes.StatusCode, message string, data any) Res {
	return Res{BaseResponse{status: status, statusCode: statusCode, message: message, data: data}}
}
