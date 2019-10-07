// req_test
package httpreqest

import (
	"fmt"
	"testing"
)

func Test001(test *testing.T) {

	ss := Reqest("get", "http://localhost:7004/test", "", 30)
	// s := bufio.NewScanner(ss.Response_data).Text()

	// fmt.Println(s)
	fmt.Println(ss)
	fmt.Println(ss.End_time - ss.Start_time)
}
