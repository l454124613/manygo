// test_content
package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type time_st struct {
	num_sec int //运行的时间
	num_men int //变化的人数
	status  int //状态1：平均递增；2：保持不变；3：平均递减
}

type time_run struct {
	num_sec int //运行的时间
	num_men int //运行的人数
	status  int //状态1：增；2：保持不变；3：减
}

func main() {
	var time_list []time_st
	// time_list = append(time_list, time_st{status: 1, num_sec: 5, num_men: 13})
	// time_list = append(time_list, time_st{status: 1, num_sec: 13, num_men: 5})
	// time_list = append(time_list, time_st{status: 1, num_sec: 5, num_men: 5})
	time_list = append(time_list, time_st{status: 1, num_sec: 1, num_men: 2300})
	time_list = append(time_list, time_st{status: 2, num_sec: 35, num_men: 0})
	time_list = append(time_list, time_st{status: 3, num_sec: 1, num_men: 2300})
	// time_list = append(time_list, time_st{status: 1, num_sec: 5, num_men: 5})
	// time_list = append(time_list, time_st{status: 3, num_sec: 13, num_men: 5})
	// time_list = append(time_list, time_st{status: 3, num_sec: 5, num_men: 5})
	var run_list []time_run
	// now_men := 0
	for _, time_st_1 := range time_list {
		switch time_st_1.status {
		case 1:
			// fmt.Println(time_st_1)
			if time_st_1.num_men > time_st_1.num_sec {
				ts := time_st_1.num_men / time_st_1.num_sec
				is_more := time_st_1.num_men % time_st_1.num_sec
				real_num := 0
				for i := 0; i < time_st_1.num_sec; i++ {
					i2 := i + 1
					if (is_more*i2/time_st_1.num_sec - real_num) < 1 {
						run_list = append(run_list, time_run{status: 1, num_sec: 1, num_men: ts})
					} else {
						run_list = append(run_list, time_run{status: 1, num_sec: 1, num_men: (ts + 1)})
						real_num++
					}
				}

			} else if time_st_1.num_men < time_st_1.num_sec {
				ts := time_st_1.num_sec / time_st_1.num_men
				is_more := time_st_1.num_sec % time_st_1.num_men
				real_num := 0
				for i := 0; i < time_st_1.num_men; i++ {
					i2 := i + 1
					if (is_more*i2/time_st_1.num_men - real_num) < 1 {
						run_list = append(run_list, time_run{status: 1, num_sec: ts, num_men: 1})
					} else {
						run_list = append(run_list, time_run{status: 1, num_sec: (ts + 1), num_men: 1})
						real_num++
					}
				}

			} else {
				for i := 0; i < time_st_1.num_men; i++ {
					run_list = append(run_list, time_run{status: 1, num_sec: 1, num_men: 1})

				}

			}

		case 2:
			// fmt.Println(time_st_1)
			run_list = append(run_list, time_run{status: 2, num_sec: time_st_1.num_sec, num_men: 0})
		case 3:
			// fmt.Println(time_st_1)
			if time_st_1.num_men > time_st_1.num_sec {
				ts := time_st_1.num_men / time_st_1.num_sec
				is_more := time_st_1.num_men % time_st_1.num_sec
				real_num := 0
				for i := 0; i < time_st_1.num_sec; i++ {
					i2 := i + 1
					if (is_more*i2/time_st_1.num_sec - real_num) < 1 {
						run_list = append(run_list, time_run{status: 3, num_sec: 1, num_men: ts})
					} else {
						run_list = append(run_list, time_run{status: 3, num_sec: 1, num_men: (ts + 1)})
						real_num++
					}
				}

			} else if time_st_1.num_men < time_st_1.num_sec {
				ts := time_st_1.num_sec / time_st_1.num_men
				is_more := time_st_1.num_sec % time_st_1.num_men
				real_num := 0
				for i := 0; i < time_st_1.num_men; i++ {
					i2 := i + 1
					if (is_more*i2/time_st_1.num_men - real_num) < 1 {
						run_list = append(run_list, time_run{status: 3, num_sec: ts, num_men: 1})
					} else {
						run_list = append(run_list, time_run{status: 3, num_sec: (ts + 1), num_men: 1})
						real_num++
					}
				}

			} else {
				for i := 0; i < time_st_1.num_men; i++ {
					run_list = append(run_list, time_run{status: 3, num_sec: 1, num_men: 1})

				}

			}
		default:
			fmt.Print("计算时间序列发生错误1001")
			return

		}

	}
	// fmt.Println(run_list)
	now_men := 0
	// now_time := 0
	var context_cancel_list []context.CancelFunc
	for _, run_l := range run_list {
		switch run_l.status {
		case 1:
			// now_men += run_l.num_men
			for i := 0; i < run_l.num_men; i++ {
				now_men++
				ctxa, cancel := context.WithCancel(context.Background())
				context_cancel_list = append(context_cancel_list, cancel)
				go work(ctxa, "work"+strconv.Itoa(now_men))
			}
			for i := 0; i < run_l.num_sec; i++ {

				time.Sleep(time.Second)

				// print_line(now_men)
			}

		case 2:
			for i := 0; i < run_l.num_sec; i++ {
				// now_men += run_l.num_men
				time.Sleep(time.Second)

			}
		case 3:
			for i := 0; i < run_l.num_men; i++ {
				now_men--

				context_cancel_list[now_men]()

				context_cancel_list = context_cancel_list[:now_men]

			}
			for i := 0; i < run_l.num_sec; i++ {

				time.Sleep(time.Second)

			}
		default:
			fmt.Print("显示时间序列发生错误1002")
			return

		}
	}

	// go work(ctxa, "work1")
	// time.Sleep(time.Second)
	// cancel()
	// time.Sleep(3 * time.Second)
	time.Sleep(5 * time.Second)
}

// func print_line(num int) {
// 	for i := 0; i < num; i++ {
// 		fmt.Print("|")
// 	}
// 	fmt.Println("                                                                      -", num)
// }

func work(ctx context.Context, name string) {
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
