package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"unsafe"
)

// testServerFunc returns the server the proxy handle will be proxying for test purpose
var testServerFunc = func(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var (
			b    bytes.Buffer
			_, _ = b.WriteString("Made it to the server")
		)

		rw.Write(b.Bytes())
	}
}

func TestNewTwitterRequest(t *testing.T) {
	var (
		tc, _  = NewTwitterClient(false)
		server = httptest.NewServer(testServerFunc(t))
		// url, _ = url.Parse(server.URL)
		req *http.Request
	)

	defer server.Close()

	req, err := tc.NewTwitterRequest()
	if err != nil {
		t.Fatalf("Incorrect size of %v", req)
	}

	if unsafe.Sizeof(req) == 0 {
		t.Fatalf("Incorrect size of %v", req)
	}
}

func TestGetHTTPResponse(t *testing.T) {
	var (
		tc, _  = NewTwitterClient(false)
		server = httptest.NewServer(testServerFunc(t))
		req    *http.Request

		b    bytes.Buffer
		_, _ = b.WriteString("Made it to the server")
	)

	res, err := tc.GetHTTPResponse(req)
	if err != nil {
		t.Fatalf("GetHTTPResponse should of not passed nil")
	}

	ok := bytes.Equal(b, res)
	if !ok {
		t.Fatalf("Failed to get the correct response")
	}
}
