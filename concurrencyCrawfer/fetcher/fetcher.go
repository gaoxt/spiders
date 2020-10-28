package fetcher

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchPost(url string, postData []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "wx.laomassf.com")
	req.Header.Set("Referer", "https://servicewechat.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.10(0x17000a21) NetType/WIFI Language/zh_CN")

	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
