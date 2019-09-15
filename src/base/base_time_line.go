// base_time_line
package base

import (
	"fmt"
)

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Time_st struct {
	num_sec int //运行的时间
	num_men int //变化的人数
	status  int //状态1：平均递增；2：保持不变；3：平均递减
}

type time_run struct {
	num_sec int //运行的时间
	num_men int //运行的人数
	status  int //状态1：增；2：保持不变；3：减
}

var time_list []time_st //存放时间对应的序列
var run_list []time_run //存放执行对应的时间序列

//方法用户添加时间计划
func Add_time_list(st Time_st) string {
	if st.num_sec == nil || st.num_sec <= 0 {
		st.num_sec = 1
	}
	if st.num_men == nil || st.num_men <= 0 {
		st.num_men = 1
	}
	if st.status == nil || st.status != 1 || st.status != 2 || st.status != 3 {
		return "时间方式不正确，错误：1004"
	}
	time_list = append(time_list, st)
	return nil
}

//方法用户计算时间计划对应的运行计划
func Calc_time_run() string {
	if len(time_list) == 0 {
		return "没有时间计划，无法计算，错误：1003"
	}
	for _, time_st_1 := range time_list {
		switch time_st_1.status {
		case 1:
			//增加人员
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
			// 人员保持不变
			run_list = append(run_list, time_run{status: 2, num_sec: time_st_1.num_sec, num_men: 0})
		case 3:
			// 人员减少
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

			return "计算时间序列发生错误,错误：1001"

		}

	}
	return nil

}

type run_func func(ctx context.Context, now_men int)

func Xxx(fun run_func) string {
	now_men := 0
	var context_cancel_list []context.CancelFunc
	for _, run_l := range run_list {
		switch run_l.status {
		case 1:
			for i := 0; i < run_l.num_men; i++ {
				now_men++
				ctxa, cancel := context.WithCancel(context.Background())
				context_cancel_list = append(context_cancel_list, cancel)
				// go work(ctxa, "work"+strconv.Itoa(now_men))
				go fun(ctxa, now_men)
			}
			for i := 0; i < run_l.num_sec; i++ {

				time.Sleep(time.Second)
			}

		case 2:
			for i := 0; i < run_l.num_sec; i++ {
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
			return "显示运行时间序列发生错误，1002"

		}
	}
	return nil
}
