package pighosts

import (
	"regexp"
	"strings"
)

func isSpecificHost(s string) bool {
	for i := range filterSpecificHostTmp {
		if filterSpecificHostTmp[i] == s {
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

func prepareHostsList(downloadHosts []string) map[string]int {
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

	return hosts
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
