package internalhttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  interface{}
	Error string
}

func (r *Response) setError(err error) {
	r.Error = err.Error()
}

func (r *Response) setData(data interface{}) {
	r.Data = data
}

func respondWithJSON(w http.ResponseWriter, code int, res Response) {
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func (r *Response) JSON(w http.ResponseWriter, code int) {
	response, _ := json.Marshal(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
