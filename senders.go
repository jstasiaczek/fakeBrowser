package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func WriteFile(url string) error {
	md5 := GetUniqueId()
	return ioutil.WriteFile(App.DirectoryToSave+"/"+md5+".txt", []byte(url), 0777)
}

func SendUrl(urlToOpen string) error {
	form := url.Values{"url": {urlToOpen}}
	req, err := http.NewRequest("POST", App.HttpUrl+"/openUrl", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("secret", App.SecretKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}
	_, err = client.Do(req)
	return err
}
