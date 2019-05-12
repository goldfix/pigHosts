package pighosts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// NON_ROUTABLE 0.0.0.0
var NON_ROUTABLE string = "0.0.0.0"

// LOCALHOST 127.0.0.1
var LOCALHOST string = "127.0.0.1"

// removeLocalHost
func removeLocalHost(s string) (string, error) {
	s = strings.ReplaceAll(s, "127.0.0.1", "")
	s = strings.ReplaceAll(s, "0.0.0.0", "")
	s = strings.TrimSpace(s)
	return s, nil
}

// removeComments
func removeComments(s string) (string, error) {
	pos := strings.Index(s, "#")
	if pos > -1 {
		s = s[0:pos]
	}
	s = strings.TrimSpace(s)
	return s, nil
}

// getRemoteList
func getRemoteList(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Status different 200 (%s, %d)", resp.Status, resp.StatusCode)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	f := func(c rune) bool {
		return c == '\n'
	}

	r := strings.FieldsFunc(strings.ReplaceAll(string(b), "\r\n", "\n"), f)
	return r, nil
}

// prepareHostsList
func prepareHostsList(urls []string) (map[string]int, error) {
	hosts := make(map[string]int, 0)

	for u := range urls {
		lstHost, err := getRemoteList(urls[u])
		if chkErr(err) {
			return nil, err
		}

		for l := range lstHost {
			hst := lstHost[l]
			hst, err = removeComments(hst)
			if chkErr(err) {
				return nil, err
			}
			hst, err = removeLocalHost(hst)
			if chkErr(err) {
				return nil, err
			}
		}
	}
	return hosts, nil
}

//prepareHostFile
func prepareHostFile(hosts map[string]int) ([]string, error) {
	result := make([]string, 0)
	for k := range hosts {
		if k != "" {
			result = append(result, k)
		}

	}

	return nil, nil
}

//writeHostFile
func writeHostFile() error {

	return nil
}

//chkErr
func chkErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
