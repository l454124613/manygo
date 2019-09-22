// req
package httpreqest

import (
	// "fmt"
	"bytes"
	"net/http"
	"strings"
	"time"
)

func Reqest(method string, url string, data string, time_out int) (int, http.Header, string, int64, int64) {
	client := &http.Client{Timeout: time.Duration(time_out) * time.Second}
	var reqest *http.Request
	var err error
	if data == "" {
		reqest, err = http.NewRequest(method, url, nil)
		if err != nil {
			return -1, nil, "请求错误：" + err.Error(), time.Now().UnixNano(), time.Now().UnixNano()
		}
	} else {
		reqest, err = http.NewRequest(method, url, strings.NewReader(data))
		if err != nil {
			return -1, nil, "请求错误：" + err.Error(), time.Now().UnixNano() / 10000, time.Now().UnixNano() / 10000
		}
		defer reqest.Body.Close()
	}

	st := time.Now().UnixNano() / 10000
	response, err2 := client.Do(reqest)
	if err2 != nil {
		return -1, nil, "响应错误：" + err2.Error(), st, time.Now().UnixNano() / 10000
	} else {
		defer response.Body.Close()
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	s := buf.String()
	return response.StatusCode, response.Header, s, st, time.Now().UnixNano() / 10000

}
