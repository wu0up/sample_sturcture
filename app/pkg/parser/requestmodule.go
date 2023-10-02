package parser

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

var client = &http.Client{}

func GetHttpClientSession() *http.Client {
	return client
}

// HttpProcessHeader 通用的 HTTP 请求处理函数
func HttpProcessHeader(method, url string, payload *strings.Reader, reqHeader map[string]string) ([]byte, error) {
	// client := &http.Client{}

	fmt.Println("method:", method, "url", url, payload)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("HttpProcessHeader ERROR:", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	for k, v := range reqHeader {
		req.Header.Add(k, v)
	}
	fmt.Println("req.Header:", req.Header)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("err:", err, "res", res, "res.body", res.Body)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// 若StatusCode非2开头(如400)，则返回error并结束流程
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", res.StatusCode)
	}

	return body, nil
}

func intsToString(ints []int) string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ",")
}

func getlocationInfo(locationIDs []int, host string, reqHeader map[string]string) ([]map[string]interface{}, error) {
	locationIDList := make([]string, len(locationIDs))

	for i, locationID := range locationIDs {
		locationIDList[i] = strconv.Itoa(locationID)
	}
	url := fmt.Sprintf("%s/json/api-locations_table?location_ids=[%s]", host, strings.Join(locationIDList, ","))
	// reqHeader := map[string]string{}
	body, err := HttpProcessHeader("GET", url, nil, reqHeader)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		if items, ok := data["items"].([]interface{}); ok {
			locationInfo := []map[string]interface{}{}
			for _, item := range items {
				if location, ok := item.(map[string]interface{}); ok {
					locationInfo = append(locationInfo, location)
				}
			}
			return locationInfo, nil
		}
	}

	return nil, fmt.Errorf("Failed to parse location data")
}

func login(host, userName, password string) (*http.Client, error) {
	password = strings.ReplaceAll(password, "+", "%2B")
	password = url.QueryEscape(password)
	fmt.Println("password:", password)
	loginURL := fmt.Sprintf("%s/__login__?username=%s&password=%s", host, userName, password)
	fmt.Println("loginURL:", loginURL)
	reqHeader := map[string]string{}

	body, err := HttpProcessHeader("GET", loginURL, nil, reqHeader)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if msg, ok := response["msg"].(string); ok && msg == "ok" {
		cookieHeader := response["Set-Cookie"].(string)
		sessionID := strings.Split(cookieHeader, ";")[0]

		// Create a new http.Client with the session ID in the cookie
		client := &http.Client{}
		client.Jar, _ = cookiejar.New(nil)
		cookies := []*http.Cookie{
			{
				Name:  "session_cookie_name", // Set the appropriate cookie name
				Value: sessionID,
			},
		}
		url, _ := url.Parse(host)
		client.Jar.SetCookies(url, cookies)

		return client, nil
	}

	return nil, fmt.Errorf("Login failed")
}

// func main() {
// 	host := "https://xintop.viewlinc.tw"
// 	userName := "api"
// 	password := "&kHk1La*TM6#"
// 	locationIDs := map[int]interface{}{
// 		1236: "幸託機房濕度",
// 		1235: "XinTop-DC",
// 	}

// 	sessionID, err := login(host, userName, password)
// 	if err != nil {
// 		fmt.Println("Login error:", err)
// 		return
// 	}

// 	client := &http.Client{}
// 	locationInfo, err := getlocationInfo(client, locationIDs, host, sessionID)
// 	if err != nil {
// 		fmt.Println("Get location info error:", err)
// 		return
// 	}

// 	fmt.Println("Location Info:", locationInfo)
// }
