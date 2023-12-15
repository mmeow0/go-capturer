package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func Create(address string, token string, data []byte) {
	apiUrl := fmt.Sprintf("http://%s", address)
	resource := "/packet"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/packet"

	bodyReader := bytes.NewReader(data)

	req, err := http.NewRequest("POST", urlStr, bodyReader)
	if err != nil {
		log.Warnln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Warnln(err)
	}

	defer response.Body.Close()

}
