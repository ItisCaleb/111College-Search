package routes

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var Limiters = map[string]*rate.Limiter{}

func Root(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./index.html")
}

func Data(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./111.json")
}

func SearchSchool(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		ip, _, _:= net.SplitHostPort(r.RemoteAddr)
		if limiter,ok := Limiters[ip];ok {
			if allow := limiter.AllowN(time.Now(),5);!allow{
				rw.WriteHeader(http.StatusTooManyRequests)
				fmt.Fprintf(rw,"星爆爆")
				return
			}
		}else {
			Limiters[ip] = rate.NewLimiter(1,150)
		}
		data,_ := ioutil.ReadAll(r.Body)
		body,err := simplejson.NewJson(data)
		if err != nil {
			return
		}
		s := Search(body.Get("name").MustString(),body.Get("department").MustString())
		fmt.Fprintf(rw, string(s))
	}
}