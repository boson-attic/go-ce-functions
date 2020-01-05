package function

import (
	"net/http"
)

func Handle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("OK\n"))
}
