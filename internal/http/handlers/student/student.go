package student

import "net/http"

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hellow this is great api "))
	}
}