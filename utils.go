package pighosts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var NumHostPerLine = 9

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
func removeLocalHost(s string) string {

	if !isSpecificHost(s) {
		s = strings.ReplaceAll(s, LocalHost, "")
		s = strings.ReplaceAll(s, NonRoutable, "")
		return strings.TrimSpace(s)
	}
	return ""

}

// removeComments
func removeComments(s string) string {
	pos := strings.Index(s, "#")
	if pos > -1 {
		s = s[0:pos]
	}
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
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
		if ChkErr(err) {
			return nil, err
		}

		for l := range lstHost {
			hst := lstHost[l]
			hst = removeComments(hst)
			hst = removeLocalHost(hst)
			if hst == "" {
				continue
			}
			hosts[hst]++
		}
	}

	return hosts, nil
}

func splitHostPerLine(hosts []string) ([]string, error) {
	hosts := make([]string, 0)

	for k, v := range hosts {

		for index := 0; index < NumHostPerLine; index++ {

		}
		hosts
	}
	return nil, nil
}

// ChkErr check returned error
func ChkErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
