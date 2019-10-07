// util_test.go
package util

import (
	. "fmt"
	"testing"
)

func Test001(test *testing.T) {

	Println("11")
}

func Test002(test *testing.T) {
	Println(GetHostName())
	Println(GetOutboundIP())
	Println(GetSys())

}
