package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServeHTTP_callsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	handler := &testHandler{}

	route := Route{"", "", "", handler}
	route.ServeHTTP(w, r)

	assert.Equal(t, handler.w, w)
	assert.Equal(t, handler.r, r)
}

type testHandler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.w = w
	h.r = r
}
