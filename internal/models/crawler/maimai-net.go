package crawler

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

// Run starts the crawler
var (
	Sid      string
	Password string
)

type crawlerClient struct {
	client *http.Client
}

func NewCrawlerClient() *crawlerClient {
	jar, _ := cookiejar.New(nil)
	return &crawlerClient{
		client: &http.Client{
			Jar: jar,
		},
	}
}

func (c *crawlerClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

// PostForm 發送POST請求
func (c *crawlerClient) PostForm(url string, data map[string][]string) (*http.Response, error) {
	return c.client.PostForm(url, data)
}

func Run(friendid string) {
	// getProfile(friendid)
}

func login() crawlerClient {
	cli := NewCrawlerClient()

	resp, err := cli.PostForm("https://maimaidx-eng.com/maimai-mobile/login/loginProcess/", map[string][]string{
		"sid":      {Sid},
		"password": {Password},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return *cli
}

func getProfile(friendid string) http.Response {
	cli := NewCrawlerClient()

	resp, err := cli.Get("https://maimaidx-eng.com/maimai-mobile/friend/profile/?friendId=" + friendid)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return *resp
}
