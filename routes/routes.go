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
				ServerError(rw,http.StatusTooManyRequests,"那你也是很快ㄟ")
				return
			}
		}else {
			Limiters[ip] = rate.NewLimiter(1,150)
		}
		data,_ := ioutil.ReadAll(r.Body)
		body,err := simplejson.NewJson(data)
		name, err := body.Get("name").String()
		department, err := body.Get("department").String()
		subject, err := body.Get("subject").Map()
		if err != nil {
			ServerError(rw,400,"屁眼")
			return
		}
		if len([]rune(name)) < 40 && len([]rune(department)) < 40{
			s := Search(name,department,subject)
			fmt.Fprintf(rw, string(s))
		}else {
			ServerError(rw,400,"沒有名字那麼長的學校好ㄇ")
		}

	}
}

func ServerError(rw http.ResponseWriter, status int, str string){
	rw.WriteHeader(status)
	fmt.Fprintf(rw,str)
}