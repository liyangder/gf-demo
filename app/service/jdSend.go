package service

import (
	"github.com/gogf/gf/frame/g"
)

var Jd = JdSend{}

type JdSend struct {
}

func (*JdSend) Test() {

	c := g.Client()
	c.SetCookie("sid", "ub3n04dqktficurme7nph5lrl0")
	//c.SetCookie('', "lang=zh-cn")

	response, err := c.Post("http://local.uss2.com/index.php?m=location&f=getError", g.Map{
		"page":      1,
		"page_size": 2,
	})

	if err != nil {
		panic(err)
	} else {
		defer response.Close()
		response.RawDump()

		//fmt.Println(response.ReadAllString())
	}

}
