// test_base_time_line
package main

import (
	"base"
	"fmt"
	"time"
)

func main() {
	base.Add_time_list(base.Time_st{status: 1, num_sec: 5, num_men: 13})
	base.Add_time_list(base.Time_st{status: 2, num_sec: 15, num_men: 13})
	base.Add_time_list(base.Time_st{status: 3, num_sec: 5, num_men: 13})
	base.Calc_time_run()
	base.Xxx(work)
	time.Sleep(5 * time.Second)

}
func work(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " get message to quit")
			return
		default:
			fmt.Println(name, " is running")
			time.Sleep(time.Second)
		}
	}
}
