package concurrent

import (
	"fmt"
	"testing"
	"time"
)

// func main() {
// 	base.Add_time_list(base.Time_st{status: 1, num_sec: 5, num_men: 13})
// 	base.Add_time_list(base.Time_st{status: 2, num_sec: 15, num_men: 13})
// 	base.Add_time_list(base.Time_st{status: 3, num_sec: 5, num_men: 13})
// 	base.Calc_time_run()
// 	base.Xxx(work)
// 	time.Sleep(5 * time.Second)

// }
func Test001(test *testing.T) {

	fmt.Println("11")
}
func Test002(test *testing.T) {
	Add_time_list(Time_st{status: 1, num_sec: 2, num_men: 2300})
	Add_time_list(Time_st{status: 2, num_sec: 2, num_men: 13000})
	Add_time_list(Time_st{status: 3, num_sec: 2, num_men: 1300})
	// Calc_time_run()
	Add_params("kk", []string{"0", "1", "2", "3"})
	Add_params("dd", []string{"5", "6", "7"})
	Add_params("cc", []string{"a", "b", "c", "d", "e"})

	Run(run_func, end_func, 30)
	// time.Sleep(5 * time.Second)
}

func run_func(num Input_type) {
	fmt.Println(num.men, "running", num.paras["kk"], num.paras["dd"], num.paras["cc"])
	time.Sleep(2 * time.Second)
}

func end_func(num Input_type) {
	fmt.Println(num.men, "stop", num.paras["kk"], num.paras["dd"])
	// time.Sleep(time.Second)
}
