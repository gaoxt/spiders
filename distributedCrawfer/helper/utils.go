package helper

import (
	"encoding/json"
	"strconv"
	"time"
)

func CreateDateFormat(createDate string) string {
	i, _ := strconv.ParseInt(createDate[6:len(createDate)-5], 10, 64)
	tm := time.Unix(i, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func UrlPathFormat(urlPath string) string {
	return "https://wx.laomassf.com" + urlPath
}

func ParserData(data interface{}) map[string]interface{} {
	var i interface{}
	json.Unmarshal([]byte(data.(string)), &i)
	jData, _ := i.(map[string]interface{})
	return jData
}
