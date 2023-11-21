package crawler

import (
	models "discordbot/internal/models/browser"
	"log"
	"time"

	"github.com/go-rod/rod/lib/proto"
)

// Run starts the crawler
var (
	Sid      string
	Password string
	cli      = models.NewPage("https://maimaidx-eng.com/")
)

func Run(friendid string, sid string, password string) []byte {
	Sid = sid
	Password = password
	save, err := getProfile(friendid)

	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	return save
}

func login() {
	cli.Page.MustElement("dt").MustClick().MustWaitStable()

	cli.Page.MustElement("input[name='sid']").MustInput(Sid)
	cli.Page.MustElement("input[name='password']").MustInput(Password)
	err := cli.Page.MustElement("input[type='submit']").Click(proto.InputMouseButtonLeft, 1)
	if err != nil {
		log.Printf("click login button error: %s", err.Error())
	} else {
		cli.Page.WaitLoad()
		time.Sleep(time.Second * 5)
	}
}

func getProfile(friendid string) ([]byte, error) {
	defer cli.Close()
	login()
	cli.Page.MustNavigate(`https://maimaidx-eng.com/maimai-mobile/friend/search/searchUser/?friendCode=` + friendid).MustWaitStable()

	// time.Sleep(time.Second * 5)
	targetDiv := cli.Page.MustElement("body > div.wrapper.main_wrapper.t_c > div.see_through_block.m_15.m_t_5.p_10.t_l.f_0.p_r > div.basic_block.p_10.f_0")
	_ = targetDiv.MustScreenshot("data/" + friendid + ".png")
	// _ = targetDiv.MustScreenshot()
	return []byte(targetDiv.MustScreenshot()), nil
}
