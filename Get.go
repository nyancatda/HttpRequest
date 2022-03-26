/*
 * @Author: NyanCatda
 * @Date: 2022-02-09 20:30:52
 * @LastEditTime: 2022-03-27 02:51:55
 * @LastEditors: NyanCatda
 * @Description: Get请求方法封装
 * @FilePath: \HttpRequest\Get.go
 */
package HttpRequest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * @description: GET请求封装
 * @param {string} url 请求地址
 * @param {[]string} Header 请求头
 * @return {[]byte} 返回内容
 * @return {*http.Response} 请求响应信息
 * @return {error} Error
 */
func GetRequest(URL string, Header []string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, nil, err
	}

	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		// 如果解析失败则不设置请求头
		if len(Headervalue) <= 0 {
			return nil, nil, errors.New("Header Error")
		}
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp, err
}
