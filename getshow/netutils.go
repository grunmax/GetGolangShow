package getshow

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

const HTTP_WAIT = 600

func GetHttpBytes(url string) []byte {
	timeout := time.Duration(HTTP_WAIT * time.Second)
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&options)
	request, _ := http.NewRequest("GET", url, nil)
	client := http.Client{
		Jar:     jar,
		Timeout: timeout,
	}
	response, e := client.Do(request)
	Err("No http response", e)
	if response.StatusCode != 200 {
		Err(fmt.Sprintf("Http response for: %s", url), response.StatusCode)
	}
	body := response.Body
	defer body.Close()
	content, _ := ioutil.ReadAll(body)
	return content
}
