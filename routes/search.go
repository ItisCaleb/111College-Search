package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"time"
)

type School struct {
	Id       int               `json:"序號"`
	Num      string            `json:"校碼"`
	Name     string            `json:"學校名稱"`
	Depart   string            `json:"系組名稱"`
	Standard map[string]string `json:"學測檢定標準"`
	English  string 		   `json:"英聽檢定標準"`
	Subject  map[string]string `json:"採計科目及加權"`
	Priority map[string]string `json:"同分參酌"`
	Tech     string			   `json:"術科類別"`
	Desc     string            `json:"選系說明"`
}

var colleges []School

func init() {
	data, _ := ioutil.ReadFile("./111.json")
	json.Unmarshal(data, &colleges)
}

func Search(name string, department string, subject map[string]interface{}) []byte {
	var r []School
	var sub []string
	for k, v := range subject{
		if reflect.TypeOf(v).Kind() == reflect.Bool{
			if v == true {
				sub = append(sub,k)
			}
		}
	}
	n := []rune(name); d := []rune(department)
	defer timeTrack(time.Now(), "Search")
	for _, college := range colleges {
		if SearchCollege(n,d,sub,college) {
			r = append(r, college)
		}
	}
	data, _ := json.Marshal(r)
	return data
}


func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s at %s\n", name, elapsed,time.Now().Format("2006/01/02;15:04"))
}

func SearchCollege(name []rune, department []rune, subject []string, college School) bool {
	lname := len(name)
	ldepartment := len(department)
	cname := []rune(college.Name)
	cdepart := []rune(college.Depart)
	c1 := 0; c2 := 0
	for i := 0; i < len(cname) && c1 < lname; i++{
		if name[c1] == cname[i]{
			c1++
		}
	}
	for i := 0; i < len(cdepart) && c2 < ldepartment; i++{
		if department[c2] == cdepart[i]{
			c2++
		}
	}
	for _, val := range subject{
		if _, ok := college.Subject[val];!ok{
			return false
		}
	}
	return c1 == lname && c2 == ldepartment

}


