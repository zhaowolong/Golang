package main

// from :http://sugarmanman.blog.163.com/blog/static/8107908020136713147504/
import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	go SignalProc()

	done := make(chan bool, 1)
	for {
		select {
		case <-done:
			break
		}

	}
	fmt.Println("exit")

}

func SignalProc() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGHUP, os.Interrupt)

	for {
		msg := <-sigs
		fmt.Println("Recevied signal:", msg)

		switch msg {
		default:
			fmt.Println("get sig=%v\n", msg)
		case syscall.SIGHUP:
			fmt.Println("get sighup\n")
		}
	}
}
