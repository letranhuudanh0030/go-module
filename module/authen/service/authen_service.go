package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"todo/config"
	"todo/database"
	"todo/module/authen/model"
)

func LoginApi(input model.LoginInput) (map[string]interface{}, string, bool) {
	postBody, _ := json.Marshal(map[string]string{
		"username": input.Username,
		"password": input.Password,
	})

	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", config.Get("SSO_BASE_URL")+"/api/auth/login", responseBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-csv-app-id", config.Get("SSO_APP_ID"))
	req.Header.Add("x-csv-app-type", config.Get("SSO_APP_TYPE"))

	client := &http.Client{}
	resp, error := client.Do(req)
	if error != nil {
		fmt.Println(error)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	jsonString := []byte(body)
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(jsonString), &data)

	resultData, _ := data["data"].(map[string]interface{})

	message := data["message"].(string)
	statusCode := resp.StatusCode
	status := false

	if statusCode == 200 {
		status = true
		// Save session
		store := database.Store
		time_expire := config.Get("JWT_EXPIRED_TIME")
		minutesCount, _ := strconv.Atoi(time_expire)
		store.Set(input.Username, []byte(resultData["token"].(string)), time.Duration(minutesCount)*time.Minute)
	}

	return resultData, message, status
}
