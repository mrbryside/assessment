package httputil

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	*http.Response
	err error
}

func (r *httpResponse) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
