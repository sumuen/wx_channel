package router

import (
	"bytes"
	"net/http"

	"github.com/qtgolang/SunnyNet/SunnyNet"
	sunnyHttp "github.com/qtgolang/SunnyNet/src/http"
)

// SunnyNetResponseWriter adapts SunnyNet.ConnHTTP to http.ResponseWriter
type SunnyNetResponseWriter struct {
	conn       SunnyNet.ConnHTTP
	headers    http.Header
	statusCode int
	body       bytes.Buffer
}

// NewSunnyNetResponseWriter creates a new SunnyNetResponseWriter
func NewSunnyNetResponseWriter(conn SunnyNet.ConnHTTP) *SunnyNetResponseWriter {
	return &SunnyNetResponseWriter{
		conn:       conn,
		headers:    make(http.Header),
		statusCode: http.StatusOK,
	}
}

func (w *SunnyNetResponseWriter) Header() http.Header {
	return w.headers
}

func (w *SunnyNetResponseWriter) Write(data []byte) (int, error) {
	return w.body.Write(data)
}

func (w *SunnyNetResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

// Flush sends the response to SunnyNet
func (w *SunnyNetResponseWriter) Flush() {
	// SunnyNet's StopRequest will interrupt the processing and send the response
	sHeaders := sunnyHttp.Header{}
	for k, v := range w.headers {
		for _, val := range v {
			sHeaders.Add(k, val)
		}
	}
	w.conn.StopRequest(w.statusCode, w.body.String(), sHeaders)
}

// ToStdRequest converts SunnyNet.ConnHTTP to *http.Request
func ToStdRequest(conn SunnyNet.ConnHTTP) (*http.Request, error) {
	body := conn.GetRequestBody()
	req, err := http.NewRequest(conn.Method(), conn.URL(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// Copy headers
	sunnyHeaders := conn.GetRequestHeader()
	for k, v := range sunnyHeaders {
		for _, val := range v {
			req.Header.Add(k, val)
		}
	}
	return req, nil
}
