package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func httpPost(url string, reqBody string, header map[string]string) (string, error) {
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
func httpGet(url string, reqBody string, header map[string]string) (string, error) {
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
func main() {
	url := "http://localhost:9090/jcrm-server-crm/rest/score/adjustScore"
	reqBody := `{"operCtx":{"time":1548214995980,"operator":{"namespace":"member_cloud","id":"member_cloud","fullName":"member_cloud"},"terminalId":"string","store":"1000401"},"request":{"tranId":"%s","xid":"23423423","tranTime":1548214995980,"account":{"type":"cardNum","id":"09000052"},"scoreRec":{"scoreType":"-","scoreSubject":"兑奖","score":"%s"},"scoreSource":"测试","remark":"说明 ","action":"消费","sourceCode":"测试"}}`
	rand.Seed(time.Now().Unix())
	tranId := "222230"
	score := "50"
	reqBody = fmt.Sprintf(reqBody, tranId, score)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["authorization"] = "Basic YWRtaW46d3d3LmhkMTIzLmNvbQ=="
	body, _ := httpPost(url, reqBody, header)
	fmt.Println(body)
}
