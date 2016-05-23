package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	sjson "github.com/bitly/go-simplejson"
)

//var login_url = "http://14.17.104.56:8123/sdk/callback_login/?gameid=170&platid=11"
//var login_url = "http://14.17.104.56:8123/sdk/callback_login"
var login_url = "http://14.17.104.56:8126/requestlogin"
var c = make(chan int)

func main() {
	testdata := sendSign("TESET DATA")
	_, ret := httpsend(login_url, string(testdata), "1")
	fmt.Printf(string(ret))
	fmt.Println(<-c)
}

func sendSign(data string) []byte {
	js := sjson.New()
	js.Set("data", data)
	js.Set("gameid", 170)
	js.Set("platid", 67)
	rawdata, _ := js.Encode()
	return rawdata
}

func httpsend(url, str string, count string) (bool, []byte) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(str))
	if err == nil {
		ret, _ := ioutil.ReadAll(resp.Body)
		//if err == nil {
		//	fmt.Println("resok", count)
		//}
		defer resp.Body.Close()
		return true, ret
	} else {
		fmt.Println(err, count)
		return false, []byte{}
	}
}
