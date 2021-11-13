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
	Subject  map[string]string `json:"採計科目及加權"`
	Priority map[string]string `json:"同分參酌"`
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
	var dp [40][40]int
	str1 := []rune(a)
	str2 := []rune(b)
	la := len(str1)
	lb := len(str2)
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] >= dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[la][lb] == la
}
