package routes

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
)

func Root(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./index.html")
}

func Data(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./111.json")
}

func SearchSchool(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		data,_ := ioutil.ReadAll(r.Body)
		body,err := simplejson.NewJson(data)
		if err != nil {
			return
		}
		s := Search(body.Get("name").MustString(),body.Get("department").MustString())
		fmt.Fprintf(rw, string(s))
	}
}