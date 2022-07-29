package main

//author:160team.west9B
import (
	"flag"
	"fmt"
)

var (
	url        string
	Zentaopath = "/zentao/user-login.html"
)

func init() {
	flag.StringVar(&url,
		"u",
		"null",
		"url:http://127.0.0.1/",
	)

}
func main() {
	flag.Parse()
	fmt.Println("author:160team.west9b")
	if url != "null" {
		Zentao()
		return
	}
	fmt.Println("usage_poc:./ZentaoSqli.exe -u url")
}
