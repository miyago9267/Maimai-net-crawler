package models

import (
	"log"
	"runtime"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/defaults"
)

type BrowserPage struct {
	Query   string
	Browser *rod.Browser
	Page    *rod.Page
}

func NewPage(query string) *BrowserPage {
	browser := OpenBrowser()
	page := browser.MustPage(query).MustWaitStable()
	return &BrowserPage{Query: query, Browser: browser, Page: page}
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

func (p *BrowserPage) ScrollToBottom() {
	footer := p.Page.MustElement("dt")
	footer.MustScrollIntoView()
	// p.Page.MustEval(`window.scrollTo(0, document.body.scrollHeight)`)
}

func (p *BrowserPage) Close() {
	p.Browser.Close()
}
