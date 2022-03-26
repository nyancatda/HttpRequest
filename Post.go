/*
 * @Author: NyanCatda
 * @Date: 2022-02-09 20:30:52
 * @LastEditTime: 2022-03-27 02:51:41
 * @LastEditors: NyanCatda
 * @Description: Post请求方法封装
 * @FilePath: \HttpRequest\Post.go
 */
package HttpRequest

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/**
 * @description: POST请求封装，传递Json
 * @param {string} url 请求地址
 * @param {[]string} Header 请求头
 * @param {string} requestBody 请求内容(Json)
 * @return {[]byte} 返回内容
 * @return {*http.Response} 请求响应信息
 * @return {error} Error
 */
func PostRequestJson(URL string, Header []string, requestBody string) ([]byte, *http.Response, error) {
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
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
 * @param {string} URL
 * @param {[]string} Header
 * @param {map[string]string} Data
 * @return {*}
 */
func PostRequestXWWWForm(URL string, Header []string, Data map[string]string) ([]byte, *http.Response, error) {
	urlValues := url.Values{}

	for Key, Value := range Data {
		urlValues.Add(Key, Value)
	}

	reqBody := urlValues.Encode()

	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(reqBody))
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

/**
 * @description: POST请求封装，带文件传递multipart/form-data
 * @param {string} URL 请求地址
 * @param {[]string} Header	请求头
 * @param {map[string]string} Data 请求数据
 * @param {string} FileKey 文件参数key
 * @param {[]string} FilePath 文件路径组
 * @return {*}
 */
func PostRequestFormDataFile(URL string, Header []string, Data map[string]string, FileKey string, FilePath []string) ([]byte, *http.Response, error) {
	client := http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	//写入文件
	for _, val := range FilePath {
		file, err := os.Open(val)
		if err != nil {
			return nil, nil, err
		}
		defer file.Close()

		fileWrite, err := bodyWrite.CreateFormFile(FileKey, val)
		if err != nil {
			return nil, nil, err
		}
		_, err = io.Copy(fileWrite, file)
		if err != nil {
			return nil, nil, err
		}
	}

	//写入其他参数
	for key, val := range Data {
		_ = bodyWrite.WriteField(key, val)
	}

	// 将w.w.boundary刷写到w.writer中
	bodyWrite.Close()

	// 创建请求
	req, err := http.NewRequest(http.MethodPost, URL, bodyBuf)
	if err != nil {
		return nil, nil, err
	}

	// 设置请求头
	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		// 如果解析失败则不设置请求头
		if len(Headervalue) <= 0 {
			return nil, nil, errors.New("Header Error")
		}
		req.Header.Set(Headervalue[0], Headervalue[1])
	}
	contentType := bodyWrite.FormDataContentType()
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, err
}

/**
 * @description: POST请求封装，传递multipart/form-data
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {map[string]string} Data 请求数据
 * @return {*}
 */
func PostRequestFormData(URL string, Header []string, Data map[string]string) ([]byte, *http.Response, error) {
	client := http.Client{}

	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	//写入参数
	for key, val := range Data {
		_ = bodyWrite.WriteField(key, val)
	}
	bodyWrite.Close()

	req, err := http.NewRequest(http.MethodPost, URL, bodyBuf)
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, err
}
