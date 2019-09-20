// clientServer
package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "time"
)

func Client_run() {
	srv := &http.Server{Addr: "0.0.0.0:7003"}

	// 指定路径下的处理器
	http.HandleFunc("/", chttpHandler)
	fmt.Println("0.0.0.0:7003")
	//指定监听host和port
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			fmt.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	var clo string
	for {
		// time.Sleep(time.Second)
		fmt.Printf(`Please enter "close" to close: `)
		fmt.Scanln(&clo)
		if clo == "close" {
			break
		}
		fmt.Println("")
	}
}

//http处理器
func chttpHandler(rw http.ResponseWriter, req *http.Request) {
	//获取参数
	vars := req.URL.Query()
	println("id : ", vars.Get("id"))
	println(req.URL.Path)

	//不同子路径,不同返回结果
	switch req.URL.Path {

	case "/task":
		io.WriteString(rw, "yes")
	case "/bunengshuodecaozuo":
		os.Exit(0)

	case "/favicon.ico":
		url := "/view/favicon.ico"
		http.Redirect(rw, req, url, http.StatusFound)
	default:
		io.WriteString(rw, "Hello manygo!")
	}
}
