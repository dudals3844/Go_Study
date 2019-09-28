package main

import (
	"fmt"
	//"math"
	//"log"
	//"os"
	//"time"
	//"sync"
	"io/ioutil"
	"net/http"
)

func main() {
	//Create request obj
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss",nil)
	if err != nil{
		panic(err)
	}
	req.Header.Add("User-Agent","Crawler")

	//Client obj start request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//print result
	bytes,_ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}


