package golinenotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const botBaseUrl = "https://notify-bot.line.me"

const apiBaseUrl = "https://notify-api.line.me"

type Response struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

func GetAuthorizeUrl(clientId, redirectUri, state string) string {
	url := botBaseUrl + "/oauth/authorize"
	url += "?response_type=code&scope=notify"
	url += "&client_id=" + clientId
	url += "&redirect_uri=" + redirectUri
	url += "&state=" + state
	return url
}

func GetAccessToken(clientId, clientSecret, redirectUri, code string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", botBaseUrl+"/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var bodyJSON Response
	json.Unmarshal(body, &bodyJSON)
	if bodyJSON.Status == 200 {
		return bodyJSON.AccessToken, nil
	}
	return "", errors.New(bodyJSON.Message)
}

func Send(token, message string) error {
	data := url.Values{}
	data.Set("message", message)

	req, err := http.NewRequest("POST", apiBaseUrl+"/api/notify", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var bodyJSON Response
	json.Unmarshal(body, &bodyJSON)
	if bodyJSON.Status == 200 {
		return nil
	}
	return nil
}
