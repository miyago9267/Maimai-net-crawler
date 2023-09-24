package crawler

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/go-rod/rod"
)

// Run starts the crawler
var (
	Sid      string
	Password string
)

func Run(friendid string) {
	// reqUrl := "https://maimaidx-eng.com/maimai-mobile/friend/friendDetail/?idx="

	// browser := openBrowser()
	// defer browser.MustClose()

	// page := browser.MustPage(reqUrl + friendid).MustWaitStable()

	// page.MustScreenshot(fmt.Sprintf("./data/screenshot/%s.png", friendid))
	getProfile(friendid)

}

func login() []*http.Cookie {
	client := &http.Client{}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// 添加 form 欄位
	err := w.WriteField("username", "your_username")
	if err != nil {
		log.Fatalf("Error adding form field: %v", err)
	}

	err = w.WriteField("password", "your_password")
	if err != nil {
		log.Fatalf("Error adding form field: %v", err)
	}

	req, err := http.NewRequest("POST", "https://lng-tgk-aime-gw.am-all.net/common_auth/login/sid/", &b)
	if err != nil {
		panic(err)
	}

	resq, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resq.Body.Close()

	if resq.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("login failed, %d", resq.StatusCode))
	}

	cookie := resq.Cookies()
	for _, c := range cookie {
		fmt.Println(c.Name, c.Value)
	}

	return cookie

}

func getProfile(friendid string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://maimaidx-eng.com/maimai-mobile/friend/friendDetail/?idx="+friendid, nil)
	if err != nil {
		panic(err)
	}

	cookie := login()
	for _, c := range cookie {
		req.AddCookie(c)
	}

	resq, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resq.Body.Close()

	if resq.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("login failed, %d", resq.StatusCode))
	}

	fmt.Println(resq.Body)
}

func openBrowser() *rod.Browser {
	var Browser *rod.Browser
	defer func() {
		if fatalError := recover(); fatalError != nil {
			log.Println(fatalError)
		}
	}()
	Browser = rod.New()
	Browser = Browser.Timeout(30 * time.Hour)
	err := Browser.Connect()
	if err != nil {
		return nil
	}
	return Browser
}
