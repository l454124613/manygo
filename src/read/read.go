// read
package read

import (
	// "fmt"
	"io/ioutil"
	"strings"
)

func read_self_conf() map[string]string {
	ra, _ := ioutil.ReadFile("../manygo.conf")
	var str string = string(ra[:])
	str = strings.Replace(str, "\r", "", -1)
	str_c := strings.Split(str, "\n")
	var map_str map[string]string
	map_str = make(map[string]string)
	map_str["close_time"] = "30"
	map_str["master_addr"] = "0.0.0.0:7002"
	map_str["client_addr"] = "0.0.0.0:7003"
	map_str["req_timeout"] = "30"
	for _, i := range str_c {
		i = strings.Trim(i, " ")
		if len(i) == 0 {
			continue
		}

		n := strings.Index(i, "#")
		// fmt.Println(n)
		if n > 1 || n == -1 {
			n1 := strings.Split(i, "=")
			// fmt.Println(i)
			if len(n1) == 2 && len(n1[1]) > 0 && len(n1[0]) > 0 {
				map_str[n1[0]] = n1[1]
			}

		}
	}
	return map_str
}
