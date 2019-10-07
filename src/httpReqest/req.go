// req
package httpreqest

import (
	"bytes"

	"net/http"
	"strings"
	"time"
)

type MyResData struct {
	Method           string
	Url              string
	Reqest_headers   http.Header
	Reqest_body      string
	Status_code      int
	Response_headers http.Header
	Response_data    string
	Start_time       int64
	End_time         int64
}

func Reqest(method string, url string, data string, time_out int) MyResData {
	client := &http.Client{Timeout: time.Duration(time_out) * time.Second}

	var reqest *http.Request
	var err error
	if data == "" {
		reqest, err = http.NewRequest(method, url, nil)
		reqest.Header.Add("key", "val")
		if err != nil {
			return MyResData{Method: method, Url: reqest.URL.String(), Reqest_headers: reqest.Header, Reqest_body: data, Status_code: -1, Response_data: "请求错误：" + err.Error(), Start_time: time.Now().UnixNano() / 10000, End_time: time.Now().UnixNano() / 10000}
			// string(reqest.URL), reqest.Header, reqest.Body, -1, nil, "请求错误：" + err.Error(), time.Now().UnixNano(), time.Now().UnixNano()
		}
	} else {
		reqest, err = http.NewRequest(method, url, strings.NewReader(data))
		if err != nil {
			return MyResData{Method: method, Url: reqest.URL.String(), Reqest_headers: reqest.Header, Reqest_body: data, Status_code: -1, Response_data: "请求错误：" + err.Error(), Start_time: time.Now().UnixNano() / 10000, End_time: time.Now().UnixNano() / 10000}
			// return -1, nil, "请求错误：" + err.Error(), time.Now().UnixNano() / 10000, time.Now().UnixNano() / 10000
		}
		defer reqest.Body.Close()
	}

	st := time.Now().UnixNano() / 10000
	response, err2 := client.Do(reqest)
	et := time.Now().UnixNano() / 10000
	defer response.Body.Close()
	if err2 != nil {
		return MyResData{Method: method, Url: reqest.URL.String(), Reqest_headers: reqest.Header, Reqest_body: data, Status_code: -1, Response_data: "响应错误：" + err.Error(), Start_time: st, End_time: time.Now().UnixNano() / 10000}
	} else {

		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		s := buf.String()
		return MyResData{Method: method, Url: reqest.URL.String(), Reqest_headers: reqest.Header, Reqest_body: data, Status_code: response.StatusCode, Response_headers: response.Header, Response_data: s, Start_time: st, End_time: et}
	}

}
