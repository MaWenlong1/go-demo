package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Url  string
	Body RequestBody
}
type RequestBody struct {
	Xid             string
	Tran_no         string
	Store_id        string
	Amount          int
	Fav_amount      int
	Consume_channel string
	Coupon_codes    []string
}

func readFile(filename string) (Request, error) {
	bytes, err := ioutil.ReadFile(filename)
	var res Request
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return res, err
	}
	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return res, err
	}
	return res, nil
}

func post(url, reqBody string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(body)
}

func main() {
	req, _ := readFile("request.json")
	url := req.Url
	body, _ := json.Marshal(req.Body)
	fmt.Println("请求地址：", url)
	fmt.Println("请求body:", string(body))
	fmt.Println("请求响应:", post(url, string(body)))
}
