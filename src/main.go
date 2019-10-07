// main
package main

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	h    bool
	v, V bool
	m    bool

	master string
)

func init() {
	flag.BoolVar(&h, "h", false, "帮助")

	flag.BoolVar(&v, "v", false, "查看版本")
	flag.BoolVar(&V, "V", false, "查看版本")

	flag.BoolVar(&m, "m", false, "控制或者执行")

	flag.StringVar(&master, "t", "", "输入控制端的地址包括端口号")

	flag.Usage = usage
}

func main1() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	if v || V {
		println("Version/0.0.1")
		return
	} else {
		println("1231")

	}
	// println(v)
	// println(V)
	// println(m)
	// println(master)

}

func usage() {
	fmt.Fprintf(os.Stderr, `manygo version: 0.0.1
Usage: manygo [-hvVm] [-t master_addr]

Options:
`)
	flag.PrintDefaults()
}
