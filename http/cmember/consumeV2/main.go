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

func readFile(filename string) (map[string]interface{}, error) {
	bytes, _ := ioutil.ReadFile(filename)

	var res interface{}
	json.Unmarshal(bytes, &res)

	result := res.(map[string]interface{})
	return result, nil
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
	fmt.Println("请求地址：", req["url"])
	body, _ := json.Marshal(req["body"])
	fmt.Println("请求body:", string(body))
	fmt.Println("请求响应:", post(req["url"].(string), string(body)))
}
