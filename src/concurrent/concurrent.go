package concurrent

import (
	"context"
	"fmt"
	"time"
)

type Time_st struct {
	Num_sec int //运行的时间
	Num_men int //变化的人数
	Status  int //状态1：平均递增；2：保持不变；3：平均递减
}

type time_run struct {
	num_sec int //运行的时间
	num_men int //运行的人数
	status  int //状态1：增；2：保持不变；3：减；4：停
}

var time_list []Time_st //存放时间对应的序列
var run_list []time_run //存放执行对应的时间序列

//方法用户添加时间计划
func Add_time_list(st Time_st) string {
	if st.Num_sec <= 0 {
		st.Num_sec = 1
	}
	if st.Num_men <= 0 {
		st.Num_men = 1
	}
	// if st.status != 1 || st.status != 2 || st.status != 3 {
	if st.Status < 1 || st.Status > 3 {
		return "时间方式不正确，错误：1004"
	}
	time_list = append(time_list, st)
	return calc_one_time_run(st)
}

func calc_one_time_run(time_st_1 Time_st) string {
	switch time_st_1.Status {
	case 1:
		//增加人员
		if time_st_1.Num_men > time_st_1.Num_sec {
			ts := time_st_1.Num_men / time_st_1.Num_sec
			is_more := time_st_1.Num_men % time_st_1.Num_sec
			real_num := 0
			for i := 0; i < time_st_1.Num_sec; i++ {
				i2 := i + 1
				if (is_more*i2/time_st_1.Num_sec - real_num) < 1 {
					run_list = append(run_list, time_run{status: 1, num_sec: 1, num_men: ts})
				} else {
					run_list = append(run_list, time_run{status: 1, num_sec: 1, num_men: (ts + 1)})
					real_num++
				}
			}

		} else if time_st_1.Num_men < time_st_1.Num_sec {
			ts := time_st_1.Num_sec / time_st_1.Num_men
			is_more := time_st_1.Num_sec % time_st_1.Num_men
			real_num := 0
			for i := 0; i < time_st_1.Num_men; i++ {
				i2 := i + 1
				if (is_more*i2/time_st_1.Num_men - real_num) < 1 {
					run_list = append(run_list, time_run{status: 1, num_sec: ts, num_men: 1})
				} else {
					run_list = append(run_list, time_run{status: 1, num_sec: (ts + 1), num_men: 1})
					real_num++
				}
			}

		} else {
			for i := 0; i < time_st_1.Num_men; i++ {
				run_list = append(run_list, time_run{status: 1, num_sec: 1, num_men: 1})

			}

		}

	case 2:
		// 人员保持不变
		run_list = append(run_list, time_run{status: 2, num_sec: time_st_1.Num_sec, num_men: 0})
	case 3:
		// 人员减少
		if time_st_1.Num_men > time_st_1.Num_sec {
			ts := time_st_1.Num_men / time_st_1.Num_sec
			is_more := time_st_1.Num_men % time_st_1.Num_sec
			real_num := 0
			for i := 0; i < time_st_1.Num_sec; i++ {
				i2 := i + 1
				if (is_more*i2/time_st_1.Num_sec - real_num) < 1 {
					run_list = append(run_list, time_run{status: 3, num_sec: 1, num_men: ts})
				} else {
					run_list = append(run_list, time_run{status: 3, num_sec: 1, num_men: (ts + 1)})
					real_num++
				}
			}

		} else if time_st_1.Num_men < time_st_1.Num_sec {
			ts := time_st_1.Num_sec / time_st_1.Num_men
			is_more := time_st_1.Num_sec % time_st_1.Num_men
			real_num := 0
			for i := 0; i < time_st_1.Num_men; i++ {
				i2 := i + 1
				if (is_more*i2/time_st_1.Num_men - real_num) < 1 {
					run_list = append(run_list, time_run{status: 3, num_sec: ts, num_men: 1})
				} else {
					run_list = append(run_list, time_run{status: 3, num_sec: (ts + 1), num_men: 1})
					real_num++
				}
			}

		} else {
			for i := 0; i < time_st_1.Num_men; i++ {
				run_list = append(run_list, time_run{status: 3, num_sec: 1, num_men: 1})

			}

		}
	default:

		return "计算时间序列发生错误,错误：1001"

	}
	return ""

}

//方法重新计算时间计划对应的运行计划
func Recalc_time_run() string {
	if len(time_list) == 0 {
		return "没有时间计划，无法计算，错误：1003"
	}
	if len(run_list) != 0 {
		run_list = append(run_list[0:0])
	}
	for _, time_st_1 := range time_list {
		ss := calc_one_time_run(time_st_1)
		if ss != "" {
			return ss
		}

	}
	return ""

}

//记录所有停止的编号
var all_done chan int

//作为参数应用
type Input_type struct {
	men   int               //当前运行的人数
	paras map[string]string //运行中的参数值
}

//初始化时，存参数的地方
var params []map[string]string

//添加参数的方法
func Add_params(key string, vals []string) {
	if len(params) == 0 {

		for _, i := range vals {
			mm := make(map[string]string)
			mm[key] = i
			params = append(params, mm)
		}

	} else {
		if len(params) > len(vals) {
			vals = append(vals, vals[:(len(params)-len(vals))]...)
			for n, i := range params {
				i[key] = vals[n]
			}
		} else if len(params) < len(vals) {
			params = append(params, params[:(len(vals)-len(params))]...)
			for n, i := range params {
				i[key] = vals[n]
			}
		} else {
			for n, i := range params {
				i[key] = vals[n]
			}
		}

	}

}

//并发运行的方法
func work(all_done chan int, ctx context.Context, num int, run_func func(inp Input_type), end_func func(inp Input_type), mm map[string]string) {
	for {
		select {
		case <-ctx.Done():

			if end_func == nil {
				all_done <- num
				return
			}
			end_func(Input_type{men: num, paras: mm})
			all_done <- num
			return
		default:
			run_func(Input_type{men: num, paras: mm})
			// fmt.Println(num, " is running")
			// time.Sleep(time.Second)
		}
	}
}

//进行并发
func Run(run_func func(inp Input_type), end_func func(inp Input_type), timeout_second int) string {
	run_list = append(run_list, time_run{status: 4, num_sec: 1, num_men: 1})
	now_men := 0
	all_done := make(chan int)
	var context_cancel_list []context.CancelFunc
	start_time := time.Now()
	params_lenght := len(params)
	for _, run_l := range run_list {
		switch run_l.status {
		case 1:
			for i := 0; i < run_l.num_men; i++ {
				now_men++
				ctxa, cancel := context.WithCancel(context.Background())
				context_cancel_list = append(context_cancel_list, cancel)

				if params_lenght >= now_men {
					go work(all_done, ctxa, now_men, run_func, end_func, params[now_men-1])
				} else {
					if params_lenght == 0 {
						go work(all_done, ctxa, now_men, run_func, end_func, nil)
					} else {

						line := (now_men - 1) % params_lenght
						// fmt.Println(line)
						go work(all_done, ctxa, now_men, run_func, end_func, params[line])
					}

				}

				// go work(ctxa, "work"+strconv.Itoa(now_men))

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
				go func() {
					<-all_done
				}()

			}
			for i := 0; i < run_l.num_sec; i++ {

				time.Sleep(time.Second)

			}
		case 4:

			for _, i := range context_cancel_list {
				i()
				// fmt.Println("正在通知...", n)
			}
			// close(all_done)
			ch_one := make(chan int)
			go func() {
				for i := 0; i < timeout_second; i++ {
					time.Sleep(time.Second)
				}
				ch_one <- 1
			}()
			go func(now_men int) {
				num := 0
				for {
					<-all_done
					// fmt.Println(cc, "ddd")
					num++
					// fmt.Println(num, "eee")
					if num >= now_men {
						break
					}
				}
				ch_one <- 2
			}(now_men)

			res_code := <-ch_one
			close(ch_one)
			fmt.Println("关闭代码：", res_code)
			fmt.Println(time.Since(start_time))
			return ""

		default:
			return "显示运行时间序列发生错误，1002"

		}
	}
	fmt.Println(time.Since(start_time))
	return ""
}
