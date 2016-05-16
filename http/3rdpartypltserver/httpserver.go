package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func login(rw http.ResponseWriter, req *http.Request) {
	log.Println("收到了login ip", string(req.RemoteAddr))
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("HandleHttpChan,Read Err:%s", err.Error())
		return
	}

	defer req.Body.Close()
	req.ParseForm() //解析参数，默认是不会解析的
	log.Println("id的值:", req.FormValue("para"))
	log.Println("data len ", len(data), string(data))
	io.WriteString(rw, `{"plataccount":"abcdfddd", "nickname":"如此执着"}`)
}

//更新售房信息
func pay(rw http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("HandleHttpChan,Read Err:%s", err.Error())
		return
	}

	defer req.Body.Close()
	req.ParseForm()
	log.Println("data len ", len(data), string(data))
	io.WriteString(rw, `{"plataccount":"abcdfddd", "nickname":"如此执着", "balance":10000}`)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/pay", pay)
	log.Fatal(http.ListenAndServe("127.0.0.1:8126", nil))
}
