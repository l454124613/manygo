// masterServer
package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	// "time"
)

func Server_run(addr string) {
	srv := &http.Server{Addr: addr}
	dir, _ := filepath.Abs(`.`)

	ss := http.Dir(dir + `\httpserver\view`)
	fsh := http.FileServer(ss)
	http.Handle("/view/", http.StripPrefix("/view/", fsh))
	// 指定路径下的处理器
	http.HandleFunc("/", mhttpHandler)
	fmt.Println(addr)
	//指定监听host和port
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	var clo string
	for {

		fmt.Printf(`Please enter "close" to close: `)
		fmt.Scanln(&clo)
		if clo == "close" {
			break
		}
		fmt.Println("")
	}
}

//http处理器
func mhttpHandler(rw http.ResponseWriter, req *http.Request) {
	//获取参数
	vars := req.URL.Query()
	println("id : ", vars.Get("id"))
	println(req.URL.Path)

	// println(&http.Server)
	//不同子路径,不同返回结果
	switch req.URL.Path {

	case "/heartbeat":
		io.WriteString(rw, "home")
	case "/connect":
		io.WriteString(rw, "cc")
	case "/close":
		io.WriteString(rw, "colse")
	case "/data":
		io.WriteString(rw, "data")
	case "/bunengshuodecaozuo":
		os.Exit(0)

	case "/favicon.ico":
		url := "/view/favicon.ico"
		http.Redirect(rw, req, url, http.StatusFound)
	default:
		url := "/view/index.html"
		http.Redirect(rw, req, url, http.StatusFound)
	}
}
