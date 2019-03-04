package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/lunny/csession"
)

type CheckInResp struct {
	Msg string `json:"msg"`
	Ret int    `json:"ret"`
}

func main() {
	if len(os.Args) < 4 {
		println("not enough parameters")
		return
	}
	site := os.Args[1]
	email := os.Args[2]
	passwd := os.Args[3]
	var s CheckInResp
	loginurl := site + "/auth/login"
	userurl := site + "/user"
	checkinurl := site + "/user/checkin"
	session := csession.New()
	session.HeadersFunc = func(req *http.Request) {
		csession.DefaultHeadersFunc(req)
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36")
	}
	forms := url.Values{
		"email": {email}, "passwd": {passwd},
	}

	resp1, err1 := session.PostForm(loginurl, forms)

	if err1 != nil {
		fmt.Println("login failed!")
		fmt.Println(err1)
		return
	}
	defer resp1.Body.Close()
	resp2, err2 := session.Get(userurl)

	if err2 != nil {
		fmt.Println("open user page failed!")
		fmt.Println(err2)
		return
	}
	defer resp2.Body.Close()
	resp, err3 := session.Post(checkinurl, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))

	if err3 != nil {
		fmt.Println("checkin failed")
		fmt.Println(err3)
		return
	}
	defer resp.Body.Close()
	body3, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	json.Unmarshal([]byte(body3), &s)
	fmt.Println(s.Msg)
}
