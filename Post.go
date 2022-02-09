/*
 * @Author: NyanCatda
 * @Date: 2022-02-09 20:30:52
 * @LastEditTime: 2022-02-09 20:34:13
 * @LastEditors: NyanCatda
 * @Description: Post请求方法封装
 * @FilePath: \HttpRequest\Post.go
 */
package HttpRequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
 * @description: POST请求封装，传递Json
 * @param {string} url 请求地址
 * @param {string} requestBody 请求内容(Json)
 * @param {[]string} Header 请求头
 * @return {[]byte} 返回内容
 * @return {*http.Response} 请求响应信息
 * @return {error} Error
 */
func PostRequestJson(url string, requestBody string, Header []string) ([]byte, *http.Response, error) {
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, nil, err
	}

	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp, err
}

/**
 * @description: POST请求封装，传递x-www-from-urlencoded
 * @param {string} Url
 * @param {[]string} Header
 * @param {map[string]string} Data
 * @return {*}
 */
func PostRequestXWWWForm(Url string, Header []string, Data map[string]string) ([]byte, *http.Response, error) {
	urlValues := url.Values{}

	for Key, Value := range Data {
		urlValues.Add(Key, Value)
	}

	reqBody := urlValues.Encode()

	req, err := http.NewRequest(http.MethodPost, Url, strings.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}

	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp, err
}
