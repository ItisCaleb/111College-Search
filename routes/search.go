package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func Search(name string, department string) []byte {
	var r []School
	defer timeTrack(time.Now(), "Search")
	for _, college := range colleges {
		if Lcs(name, college.Name) && Lcs(department, college.Depart) {
			r = append(r, college)
		}
	}
	data, _ := json.Marshal(r)
	return data
}


func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func Lcs(a string, b string) bool {
	str1 := []rune(a)
	str2 := []rune(b)
	la := len(str1)
	lb := len(str2)
	count := 0
	for j := 0; j < lb && count<la; j++ {
		if str1[count] == str2[j] {
			count++
		}
	}
	return count == la
}
