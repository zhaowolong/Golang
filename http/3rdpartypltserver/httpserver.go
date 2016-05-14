package main

import(
	"log"
	"net/http"
)


func login(rw http.ResponseWriter, req *http.Request){
	req.ParseForm() //解析参数，默认是不会解析的 
	log.Println("id的值:", req.FormValue("id"))
}

//更新售房信息
func pay(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm();
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/pay", pay)
	log.Fatal(http.ListenAndServe(":8126", nil))
}
