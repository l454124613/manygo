// read_test
package read

import (
	"fmt"
	"testing"
	// "time"
)

func Test001(test *testing.T) {

	fmt.Println("11")
}

func Test002(test *testing.T) {
	test1 := read_self_conf()
	fmt.Println(test1)

}
