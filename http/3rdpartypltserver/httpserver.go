package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var debug = flag.Bool("debug", true, "enable debugging")
var password = flag.String("password", "asd123", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "222.186.52.217", "the database server")
var user = flag.String("user", "sa", "the database user")

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
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")

	http.HandleFunc("/login", login)
	http.HandleFunc("/pay", pay)
	log.Fatal(http.ListenAndServe("127.0.0.1:8126", nil))
}
