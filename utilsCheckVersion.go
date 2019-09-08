package pighosts

import (
	"io/ioutil"
	"net/http"
	"time"
)

func GetVersion(currentVer string) (bool, string, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	defer client.CloseIdleConnections()
	resp, err := client.Get("https://raw.githubusercontent.com/goldfix/pigHosts/master/VERSION")
	if err != nil {
		return false, "nd", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "nd", err
	}

	result := string(body)
	if string(result) != currentVer && currentVer != "dev" {
		return true, string(result), nil
	} else {
		return false, string(result), nil
	}
}
