package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//buf := make([]byte, 1024)
	//f, _ := os.Open("/etc/profile")
	//defer f.Close()
	//for {
	//	n, _ := f.Read(buf)
	//	if n == 0 {
	//		break
	//	}
	//	os.Stdout.Write(buf)
	//}
	r, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Println(string(body))
	}
}
