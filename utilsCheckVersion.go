package pighosts

import (
	"io/ioutil"
	"net/http"
	"strings"
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

	result := strings.TrimSpace(string(body))
	currentVer = strings.TrimSpace(currentVer)

	if currentVer == "dev" || result == "nd" || result == currentVer {
		return false, string(result), nil
	} else {
		return true, string(result), nil
	}
}
