// savedata
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	. "../util"
)

var q = make(chan interface{}, 10)

func main() {
	fmt.Println("Hello World21!")
	// root_dir := `e:\testData`
	// Get_root(root_dir)
	// Add_src("test1", root_dir)

	// fmt.Println(is_e)
	ctxa, cancel := context.WithCancel(context.Background())
	go Get_detail(test, ctxa, 1)
	for i := 0; i < 100; i++ {
		Add_detail(i)
	}
	time.Sleep(2 * time.Second)
	cancel()
}

func test(aa interface{}) {
	fmt.Println(aa)
}

func Get_root(root_path string) string {
	if !Exists(root_path) {
		os.MkdirAll(root_path, 0777)
	}
	return root_path
}

func Add_timearea(name string, data []string) {

}
func Add_log(name string, data []string) {

}
func Add_err(name string, data []string) {

}

func Add_detail(v interface{}) {
	q <- v
}

func Get_detail(func_handle func(interface{}), ctx context.Context, wait_time int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			click := time.After(time.Duration(wait_time) * time.Second)
			select {
			case data := <-q:
				func_handle(data)
			case <-click:
				continue
			}
		}
	}

}

func Add_src(name string, root_dir string) string {
	if !Exists(root_dir + "/" + name) {
		os.MkdirAll(root_dir+"/"+name, 0777)
	}
	if !Exists(root_dir + "/" + name + "/timearea.csv") {
		Create_file(root_dir + "/" + name + "/timearea.csv")
	}
	if !Exists(root_dir + "/" + name + "/error.csv") {
		Create_file(root_dir + "/" + name + "/error.csv")
	}
	if !Exists(root_dir + "/" + name + "/log.csv") {
		Create_file(root_dir + "/" + name + "/log.csv")
	}
	return root_dir + "/" + name

}
