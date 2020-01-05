package function_test

import (
	"net/http/httptest"
	"testing"

	"function"
)

// TestHandle ensures that Handle accepts a valid HTTP request and responds
// with an HTTP OK.
func TestHandle(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/function", nil)
	w := httptest.NewRecorder()

	function.Handle(w, req)

	res := w.Result()
	if res.StatusCode != 200 {
		t.Fatalf("handler returned HTTP %v", res.StatusCode)
	}
}
