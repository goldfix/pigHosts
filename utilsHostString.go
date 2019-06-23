package pighosts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func isSpecificHost(s string) bool {
	for i := range specificHost {
		if specificHost[i] == s {
			return true
		}
	}
	return false
}

func removeLocalHost(s string) string {

	if !isSpecificHost(s) {
		s = strings.ReplaceAll(s, localHost, "")
		s = strings.ReplaceAll(s, nonRoutable, "")
		return strings.TrimSpace(s)
	}
	return ""

}

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

// prepareHostsList
func prepareHostsList(downloadHosts []string) (map[string]int, error) {
	hosts := make(map[string]int, 0)

	for l := range downloadHosts {
		hst := downloadHosts[l]
		hst = removeComments(hst)
		hst = removeLocalHost(hst)
		if hst == "" {
			continue
		}
		hosts[hst]++
	}

	return hosts, nil
}

func downlaodRemoteList(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Status different 200 (%s, %d)", resp.Status, resp.StatusCode)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	f := func(c rune) bool {
		return c == '\n'
	}

	r := strings.FieldsFunc(strings.ReplaceAll(string(b), "\r\n", "\n"), f)
	return r, nil
}

func splitHostPerLine(hosts map[string]int) []string {
	result := make([]string, 0)

	t := nonRoutable
	c := 0
	for v := range hosts {
		t = t + " " + v
		c++
		if c >= numHostPerLine {
			result = append(result, t)
			c = 0
			t = nonRoutable
		}
	}
	if c > 0 {
		result = append(result, t)
		c = 0
		t = nonRoutable
	}

	return result
}
