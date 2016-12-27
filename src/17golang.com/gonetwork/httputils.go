package gonetwork

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	//"log"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	//"time"
)



/**
 *   http get 方法
 *   @parm url  请求的URL
     @HTTPHEAD  请求的HTTP头
 *
 */
func httpGet(url string, httphead map[string]string) (int, string) {

	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	//reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	if httphead != nil {

		for key := range httphead {

			value := httphead[key]

			reqest.Header.Set(key, value)

		}

	}

	var bodystr string
	response, _ := client.Do(reqest)

	if response.StatusCode == 200 {

		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 16)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				bodystr += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			bodystr = string(bodyByte)
		}
	}

	defer response.Body.Close()
	return response.StatusCode, bodystr

}

/**
 * http post 基本方法
 */
func httpPostBase(paths string, httphead map[string]string, bodycontent io.Reader) (int, string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", paths, bodycontent)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//req.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//req.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//req.Header.Set("Accept-Language","zh-CN,zh;q=0.8")

	if httphead != nil {

		for key := range httphead {

			value := httphead[key]

			req.Header.Set(key, value)

		}

	}

	resp, err := client.Do(req)

	if err == nil {

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return resp.StatusCode, string(body), err
	} else {
		return -100, "", err
	}

	//fmt.Println(string(body))

	////////////////////////////////

}

/**
 *
 */
func HttpPost4String(paths string, httphead map[string]string, body string) (int, string, error) {

	return httpPostBase(paths, httphead, strings.NewReader(body))

}

/**
 * 传入参数的HTTPpost
 */
func HttpPost(paths string, httphead map[string]string, parmmap map[string]string) (int, string, error) {

	postValues := url.Values{}

	if parmmap != nil {

		for key := range parmmap {

			value := parmmap[key]
			postValues.Set(key, value)

		}
	}

	//postbody := ioutil.NopCloser(strings.NewReader(postValues.Encode()))

	postDataStr := postValues.Encode()
	postDataBytes := []byte(postDataStr)

	postBytesReader := bytes.NewReader(postDataBytes)

	return httpPostBase(paths, httphead, postBytesReader)

}

func HttpPost1() {

	resp, err := http.Post("http://api.eshowpro.com/table/remotenotify",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://api.eshowpro.com/table/remotenotify", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err

}

func GethTest() {

}