package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	max := 500
	var wg sync.WaitGroup
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func() {
			st := time.Now()
			getToken()
			elapsed := time.Since(st)
			fmt.Print(elapsed, "\t")
			wg.Done()
		}()
	}
	wg.Wait()
}

func getToken() string {
	url := "http://localhost:8080/v1/oauth/tokens?grant_type=password&username=mwl@123&password=123123&scope=read_write"
	req, _ := http.NewRequest("POST", url, strings.NewReader(""))
	// input := []byte("test_client_1:test_secret")
	// encodeString := base64.StdEncoding.EncodeToString(input)
	// req.Header.Set("authorization", "Basic "+encodeString)
	// fmt.Println(encodeString)
	req.SetBasicAuth("test_client_1", "test_secret")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return "nil"
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败", err)
		return "nil"
	}
	return string(body)

}
