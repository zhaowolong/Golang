package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	//"io"
	"io/ioutil"
	"time"
	//"log"
	//"math/rand"
	"net/http"
	"strings"
	sjson "github.com/bitly/go-simplejson"
	"git.code4.in/mobilegameserver/logging"
)

var loginUrl = "http://14.17.104.56:8000/httplogin"
var gameid = 170
var zoneid = 301

var str = fmt.Sprintf(`{"do":"plat-token-login","gameid":170,"zoneid":301,"data":{"platinfo":{"account":"","platid":0,"email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722"}}}`)
var countMsg = 0
func main(){	
	for i := 1; i < 10000; i++ {
		goindex := i 
		go connect(goindex)
	}
	c := make(chan int, 1)
	fmt.Println(<-c)
}

func connect(goindex int) {
	// plat-token-login
	plattokenlogin := fmt.Sprintf(`{"do":"plat-token-login", "gameid":170, "zoneid":301, "data":{"platinfo":{"account":"", "platid":0}}}`)
	count := fmt.Sprintf("%s: %d", "plattokenlogin", goindex)
	bOk, token := httpsend(loginUrl, plattokenlogin, count)
	if !bOk {
		logging.Error("httpsend error plat-token-login ")
		return 
	}
	js, err := sjson.NewJson(token)
	if err != nil {
		logging.Error("platt-token-login  to json error")
		return 
	}
	unigame_plat_key := js.Get("unigame_plat_key").MustString()
	unigame_plat_login := js.Get("unigame_plat_login").MustString()
	uid := js.Get("data").Get("uid").MustString()
	// select-zone
	data := "{}"
	signurl, dataSend := sendSign(uid, "request-select-zone", data, unigame_plat_key, unigame_plat_login, loginUrl, 170, 301)
	bOk, token = httpsend(signurl, string(dataSend), count)
	if !bOk {
		logging.Error("httpsend error select-zone error")
		return 
	}
	js, err = sjson.NewJson(token)
	if err != nil {
		logging.Error("select zone to json error")
		return 
	}
	gatewayurl := js.Get("data").Get("gatewayurl").MustString()
	// sendTounilight
	for j:=0; j<1000; j +=1 {
		signurl, dataSend := sendSign(uid, "Cmd.UserInfoSynRequestLbyCmd_C", "{}", unigame_plat_key, unigame_plat_login, gatewayurl, 170, 301)
		bOk, token = httpsend(signurl, string(dataSend), count)
		if !bOk {
			logging.Error("httpsend error UserInfoSynRequestLbyCmd_c")
			return 
		}
		js, err = sjson.NewJson(token)
		if err != nil {
			logging.Error("UserInfoSynRequestLbyCmd_C zone to json error")
			return 
		}
		desc := js.Get("data").Get("desc").MustString()
		countMsg += 1
		logging.Info("rev unilight%s, 第%d个携程中的第%d次访问， 共访问次数%d", desc, goindex, j, countMsg)
	}
}

func sendSign(uid, do, data, unigame_plat_key, unigame_plat_login, url string, gameid, zoneid int)(string, []byte){
	unigame_plat_timestamp := int(time.Now().Unix())
	js := sjson.New()
	js.Set("do", do)
	js.Set("data", data)
	js.Set("unigame_plat_key", unigame_plat_key)
	js.Set("unigame_plat_login", unigame_plat_login)
	js.Set("gameid", gameid)
	js.Set("zoneid", zoneid)
	js.Set("uid", uid)
	js.Set("unigame_plat_timestamp", unigame_plat_timestamp)
	rawdata,_ := js.Encode()

	hash := md5.New()
	timestr := strconv.Itoa(unigame_plat_timestamp)
	hash.Write(append(append(rawdata, ([]byte(timestr))...), unigame_plat_key...))
	sign := fmt.Sprintf("%x", hash.Sum(nil))

	signurl := fmt.Sprintf("%s?unigame_plat_sign=%s", url, sign)
	return signurl, rawdata
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
