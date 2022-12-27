package util

type err struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct{}

func newResponse() response {
	return response{}
}

func (r response) ApiError(code, message string) err {
	return err{Code: code, Message: message}
}
