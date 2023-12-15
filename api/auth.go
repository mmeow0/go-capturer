package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mmeow0/go-capturer/models"
	log "github.com/sirupsen/logrus"
)

func Login(address string, user string, password string) string {
	apiUrl := fmt.Sprintf("http://%s", address)
	resource := "/login"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/login"

	loginData := url.Values{
		"email":    {user},
		"password": {password},
	}

	resp, err := http.PostForm(urlStr, loginData)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Cannot login:", err.Error())
	}
	defer resp.Body.Close()

	loginResp := &models.LoginResponse{}
	derr := json.NewDecoder(resp.Body).Decode(loginResp)

	if derr != nil {
		log.Fatalln("Cannot decode response:", err.Error())
	}

	log.Infoln("Login success")

	return loginResp.AccessToken

}
