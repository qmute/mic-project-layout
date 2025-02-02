package ut

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
)

func Verbose(v bool) {
	if v {
		logrus.SetOutput(os.Stderr)
	} else {
		logrus.SetOutput(io.Discard) // 测试期间保持安静
	}
}

func Reader(v interface{}) io.Reader {
	b, err := json.Marshal(v)
	utee.Chk(err)
	return bytes.NewReader(b)
}

// Serve 快捷方法，把app mount到router上， 并发出请求
func Serve(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// Req 快捷方法， 生成request
func Req(method Method, path string, q url.Values, body io.Reader, contentType ...string) *http.Request {
	path = strings.TrimSpace(path)
	if path == "" {
		panic("path must be not empty")
	}

	if q != nil {
		path = path + "?" + q.Encode()
	}
	r, _ := http.NewRequest(string(method), path, body)
	if len(contentType) > 0 && contentType[0] != "" {
		r.Header.Set("Content-Type", contentType[0])
	}

	return r
}
