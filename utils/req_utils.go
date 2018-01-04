package utils

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/franela/goreq"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type errorData struct {
	Message string `json:"error_message"`
}

// GetContent ... 获取某个URL页面的内容
func GetContent(url string, headers map[string]string) (string, error) {
	req0 := goreq.Request{
		Uri:         url,
		Compression: goreq.Gzip(),
	}
	for k, v := range headers {
		req0.AddHeader(k, v)
	}
	resp0, err := req0.Do()
	if err != nil {
		return "", err
	}

	defer resp0.Body.Close()

	// check status
	if resp0.StatusCode != http.StatusOK {
		errData := errorData{}
		err = resp0.Body.FromJsonTo(&errData)
		if err != nil {
			return "", err
		}
		return "", errors.New(errData.Message)
	}

	bodyData, err := ioutil.ReadAll(resp0.Body)
	if err != nil {
		return "", err
	}
	return string(bodyData), nil
}

// GetDocument ... 获取某个URL页面的Document对象
func GetDocument(url string, headers map[string]string) (*goquery.Document, error) {
	req0 := goreq.Request{
		Uri:         url,
		Compression: goreq.Gzip(),
	}
	for k, v := range headers {
		req0.AddHeader(k, v)
	}
	resp0, err := req0.Do()
	if err != nil {
		return nil, err
	}
	defer resp0.Body.Close()
	// check status
	if resp0.StatusCode != http.StatusOK {
		errData := errorData{}
		err = resp0.Body.FromJsonTo(&errData)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errData.Message)
	}
	bodyData, err := ioutil.ReadAll(resp0.Body)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(bytes.NewReader(bodyData))
}

func DownloadFile(url, destFile string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	file, _ := os.Create(destFile)
	io.Copy(file, res.Body)
	return nil
}
