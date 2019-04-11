package main

import (
	"fmt"
	"net/http"
)

var headers map[string]string = make(map[string]string)

func init() {
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "zh-CN,zh;q=0.9"
	headers["cookie"] = "PHPSESSID=36c8n4lsbb8u63glevh1ksc9a1; webp_enabled=1; _ga=GA1.2.1167535880.1534758916; _gid=GA1.2.1330668796.1534758916; weilisessionid=aa3bf69b4f35c91ca4866315f1f300b1; wluuid=WLGEUST-02ADBA37-4B6C-DE33-2769-8697C4B575BB; wlsource=tc_pc_home; webp_enabled=0; _ga=GA1.3.1167535880.1534758916; _gid=GA1.3.1330668796.1534758916; _ba=BA0.2-20180820-51751-eyUyUL4rqUHUI1lh6uRM; qimo_seosource_e7dfc0b0-b3b6-11e7-b58e-df773034efe4=%E5%85%B6%E4%BB%96%E7%BD%91%E7%AB%99; qimo_seokeywords_e7dfc0b0-b3b6-11e7-b58e-df773034efe4=%E6%9C%AA%E7%9F%A5; accessId=e7dfc0b0-b3b6-11e7-b58e-df773034efe4; pageViewNum=1; bad_ide7dfc0b0-b3b6-11e7-b58e-df773034efe4=3c85f321-a45f-11e8-92ed-072415955da9; nice_ide7dfc0b0-b3b6-11e7-b58e-df773034efe4=3c85f322-a45f-11e8-92ed-072415955da9"
	headers["dnt"] = "1"
	headers["upgrade-insecure-requests"] = "1"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"
}

func getImageID(term string, page int) {
	println("获取图片ID.....")
	url := "https://stock.tuchong.com/api/free/search/?term=" + term + "&page=" + string(page)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Headers = headers
	respone, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	println(respone)
}

func main() {
	for i := 0; i < 7; i++ {
		getImageID("", i)
	}
	println(headers)
}
