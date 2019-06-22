package pighosts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var numHostPerLine = 9

// NonRoutable 0.0.0.0
var nonRoutable = "0.0.0.0"

// LocalHost 127.0.0.1
var localHost = "127.0.0.1"

// SpecificHost
var specificHost = []string{
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
	"localhost",
	"localhost.localdomain",
	"local",
	"broadcasthost",
	"ip6-localhost",
	"ip6-loopback",
	"ip6-localnet",
	"ip6-mcastprefix",
	"ip6-allnodes",
	"ip6-allrouters",
	"ip6-allhosts",
	"0.0.0.0",
}

// SpecificHost
var defaultHostsUrls = []string{
	"https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts",
	"https://www.squidblacklist.org/downloads/dg-ads.acl",
	"https://www.squidblacklist.org/downloads/dg-malicious.acl",
}

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

func initConf() error {
	s, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return nil
}
