package pighosts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// NonRoutable 0.0.0.0
var NonRoutable = "0.0.0.0"

// LocalHost 127.0.0.1
var LocalHost = "127.0.0.1"

// SpecificHost
var SpecificHost = []string{
	"127.0.0.1 localhost",
	"127.0.0.1 localhost.localdomain",
	"127.0.0.1 local",
	"255.255.255.255 broadcasthost",
	"::1 localhost",
	"::1 ip6-localhost",
	"::1 ip6-loopback",
	"fe80::1%lo0 localhost",
	"ff00::0 ip6-localnet",
	"ff00::0 ip6-mcastprefix",
	"ff02::1 ip6-allnodes",
	"ff02::2 ip6-allrouters",
	"ff02::3 ip6-allhosts",
	"0.0.0.0 0.0.0.0",
}

func isSpecificHost(s string) bool {
	for i := range SpecificHost {
		if SpecificHost[i] == s {
			return true
		}
	}
	return false
}

// removeLocalHost
func removeLocalHost(s string) (string, error) {

	if !isSpecificHost(s) {
		s = strings.ReplaceAll(s, LocalHost, "")
		s = strings.ReplaceAll(s, NonRoutable, "")
	}
	s = strings.TrimSpace(s)

	return s, nil
}

// removeComments
func removeComments(s string) (string, error) {
	pos := strings.Index(s, "#")
	if pos > -1 {
		s = s[0:pos]
	}
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
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
			if isSpecificHost(hst) {
				hst, err = removeLocalHost(hst)
				if chkErr(err) {
					return nil, err
				}
			}

			hosts[hst]++
		}
	}
	return hosts, nil
}

//prepareHostFile
func prepareHostFile(hosts map[string]int) ([]string, error) {
	result := make([]string, 0)
	for k := range hosts {
		if k == "" {
			continue
		}
		if strings.Index(k, LocalHost) > -1 || strings.Index(k, NonRoutable) > -1 {
			continue
		}
		if isSpecificHost(k) {
			result = append(result, k)
		} else {
			result = append(result, NonRoutable+" "+k)
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
