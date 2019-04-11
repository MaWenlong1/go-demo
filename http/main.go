package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url string, reqBody string, header map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败", err)
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}
func HttpGet(url string, reqBody string, header map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败", err)
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败", err)
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}
