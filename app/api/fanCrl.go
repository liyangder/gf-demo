package api

import (
	"fmt"
	"gf-demo/app"
	"gf-demo/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"sync"
)

var FanCrl = fanCrl{}

type fanCrl struct {
}

func (*fanCrl) Test(r *ghttp.Request) {
	ch := make(chan int, 0)
	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
		}
		defer close(ch)
	}()
	c1 := change(ch)
	c2 := change(ch)
	c3 := change(ch)

	//a := make([]int,0)
	a := make([]int, 1)

	for {
		select {
		case temp, ok := <-c1:
			app.DD(ok)
			if !ok {
				fmt.Println(111)
				c1 = nil
			} else {
				a = append(a, temp)
			}
		case temp, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				a = append(a, temp)
			}
		case temp, ok := <-c3:
			if !ok {
				c3 = nil
			} else {
				a = append(a, temp)
			}

		//	//fmt.Println(gtime.Now())
		default:
			fmt.Println(gtime.Now())
			//a, ok := <-c1
			//app.DD(a, ok)
		}
		if c1 == nil && c2 == nil && c3 == nil {
			//if c1 == nil {
			app.DD(a)
			break
		}
	}

	app.OutSuccess(r, a)
	//app.OutSuccess(r, gtime.Now().Format("Ym"))
}

func change(arr chan int) chan int {
	out := make(chan int)
	go func() {
		defer func() {
			app.DD(out)
			close(out)
		}()
		for value := range arr {
			out <- (value * value)
		}
	}()
	//close(out)
	return out
}

func (*fanCrl) WgTest(r *ghttp.Request) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 8000; i++ {
			power := grand.N(50, 188)
			price := grand.N(2, 500)
			dao.Item.Ctx(r.Context()).Data(g.Map{"price": price, "power": power}).Insert()
		}
		wg.Done()
	}()

	wg.Wait()
	app.OutSuccess(r, 1)
}
