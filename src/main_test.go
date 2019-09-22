// main_test
package main

import (
	"concurrent"
	"fmt"
	"httpReqest"

	// "read"
	"testing"
)

var ccmap []int64

func Test001(test *testing.T) {
	// ccmap = make(map[int]int64)
	concurrent.Add_time_list(concurrent.Time_st{Num_sec: 1, Num_men: 100, Status: 1})
	concurrent.Add_time_list(concurrent.Time_st{Num_sec: 299, Num_men: 100, Status: 2})
	concurrent.Run(run_func, nil, 30)
	fmt.Println(len(ccmap))
	var min int64 = 1000000000
	var max int64 = 0
	var sum int64 = 0
	for _, i := range ccmap {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
		sum += i

	}
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(sum / int64(len(ccmap)))
	// fmt.Println(ccc, ccmap)

}
func run_func(num concurrent.Input_type) {
	code, _, body, st, et := httpreqest.Reqest("get", "http://192.168.1.3:7004/test", "", 30)

	if code == 200 && len(body) > 0 {

		ccmap = append(ccmap, et-st)
	} else {
		fmt.Println(code, len(body), st, et)
	}
	// fmt.Println(code, len(body), st, et)

}
