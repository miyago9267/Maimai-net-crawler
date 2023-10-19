package models

import (
	"log"
	"runtime"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/defaults"
)

type Page struct {
	query   string
	browser *rod.Browser
	page    *rod.Page
}

func NewPage(query string) *Page {
	browser := OpenBrowser()
	page := browser.MustPage(query).MustWaitStable()
	return &Page{query: query, browser: browser, page: page}
}

func OpenBrowser() *rod.Browser {
	var Browser *rod.Browser
	os := runtime.GOOS
	defaults.Show = (os == "windows")
	// defaults.Show = true
	defer func() {
		if fatalError := recover(); fatalError != nil {
			log.Println(fatalError)
		}
	}()
	Browser = rod.New()
	Browser = Browser.Timeout(30 * time.Hour)
	err := Browser.Connect()
	if err != nil {
		panic(err)
	}
	return Browser
}

func (p *Page) Close() {
	p.browser.Close()
}

func (p *Page) FuckYouNigga() {

}
