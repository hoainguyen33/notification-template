package domain

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func FcmSendPostWithHeader(url string, data interface{}) interface{} {

	byteArray, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	str := string(byteArray)
	payload := strings.NewReader(str)

	client := &http.Client{
		Timeout: 3600 * time.Second,
	}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "key="+os.Getenv("FCM_SERVER_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	var objmap map[string]interface{}
	var objmapArray []map[string]interface{}
	err = json.Unmarshal(body, &objmap)
	err2 := json.Unmarshal(body, &objmapArray)
	if err == nil {
		return objmap
	}
	if err2 == nil {
		return objmapArray
	}

	return nil
}

func GetHtml(url string) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(body)
}
