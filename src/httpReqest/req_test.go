// req_test
package httpreqest

import (
	"fmt"
	"testing"
)

func Test001(test *testing.T) {
	fmt.Println(reqest("get", "http://192.168.1.3:7004/test", "", 30))
}
