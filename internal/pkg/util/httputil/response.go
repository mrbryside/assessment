package httputil

type response struct{}

type errResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func newResponse() response {
	return response{}
}

func (r response) ApiError(code, message string) errResp {
	return errResp{Code: code, Message: message}
}
