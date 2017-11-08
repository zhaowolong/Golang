package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"tool.lu/sandbox-server/app"
)

func main() {
	debug()
	server := app.NewApp()
	server.Run(":9090")
}

func debug() {
	go func() {
		// 这边是由于通过pprof发现问题之后，加的一段debug代码；后面会讲到
		http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
			num := strconv.FormatInt(int64(runtime.NumGoroutine()), 10)
			w.Write([]byte(num))
		})
		http.ListenAndServe("localhost:6060", nil)
	}()
}
