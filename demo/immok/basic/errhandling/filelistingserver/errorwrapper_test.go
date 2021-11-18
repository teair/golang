package filelistingserver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// errPanicError 测试返回panic错误类型
func errPanicError(writer http.ResponseWriter, r *http.Request) error {
	panic(123)
}

type testUserError string

func (e testUserError) Error() string {
	return e.Message()
}

func (e testUserError) Message() string {
	return string(e)
}

// errUserError 测试用户可见错误类型
func errUserError(writer http.ResponseWriter, r *http.Request) error {
	return testUserError("user error")
}

func errNotExists(writer http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}
func errUnknown(writer http.ResponseWriter, r *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanicError, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotExists, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// TestErrorWrapper 测试错误处理函数
func TestErrorWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h) // 模拟业务处理函数
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil,
		)
		f(response, request) // 至此数据构建完成【单独测试该函数】

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL) // 模拟真正的请求

		// 整合
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")

	if body != expectedMsg || resp.StatusCode != expectedCode {
		t.Errorf("expected:(%d,%s); got:(%d,%s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
