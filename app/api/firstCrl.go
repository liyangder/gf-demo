package api

import (
	"bytes"
	"fmt"
	"gf-demo/app/service"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"log"
	"net/url"
	"time"
)

var FirstCrl = firstCrl{}

type firstCrl struct{}

type UserInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserRequest struct {
	UserInfo *UserInfo `json:"user_info"`
}

/*
   收集器请求前: onRequest()
   收集器抓取失败:onError()
   	收集器响应后：onResponse()
   收集器收到HTML:onHTML()
   	收集器收到XML： onXML()
   	收集器抓取完后最后执行的回调：onScraped()
*/
func (*firstCrl) Test(r *ghttp.Request) {
	//urlstr := "https://news.baidu.com"
	urlstr := "https://www.qimao.com/shuku/113609/"

	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	c := colly.NewCollector()
	// 超时设定
	c.SetRequestTimeout(100 * time.Second)
	// 指定Agent信息
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", u.Host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", u.Host)
		r.Headers.Set("Referer", urlstr)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("title:", e.Text)
	})

	c.OnResponse(func(resp *colly.Response) {
		fmt.Println("response received", resp.StatusCode)

		// goquery直接读取resp.Body的内容
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))

		// 读取url再传给goquery，访问url读取内容，此处不建议使用
		// htmlDoc, err := goquery.NewDocument(resp.Request.URL.String())

		if err != nil {
			log.Fatal(err)
		}

		// 找到抓取项 <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		//htmlDoc.Find(".hotnews a").Each(func(i int, s *goquery.Selection) {
		htmlDoc.Find("body").Each(func(i int, s *goquery.Selection) {
			fmt.Println(222)
			title := s.Text()
			fmt.Printf("热点新闻 %d: %s - %s\n", i, title)
			//c.Visit(band)
		})

	})

	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
	})

	err = c.Visit(urlstr)

	fmt.Println(123)
}

// Index is a demonstration route handler for output "Hello World!".
func (*firstCrl) Index(r *ghttp.Request) {
	r.Response.Writeln("aaa")

	service.Jd.Test()

	r.Response.WriteJsonExit(gtime.New(time.Now()))

}

// Upload uploads files to ./tmp .
func (*firstCrl) Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("file_name")
	names, err := files.Save("./tmp/")
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}
