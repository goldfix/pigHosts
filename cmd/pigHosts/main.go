package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"net/http"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `
  Usage: arguments_example [-vqrh] [FILE] ...
    arguments_example (--left | --right) CORRECTION FILE
Process FILE and optionally apply correction to either left-hand side or
  right-hand side.
Arguments:
  FILE        File to process
Options:
  -h --help
  `

	arguments, _ := docopt.Parse(usage, nil, true, "", false)
	fmt.Printf("%v", reflect.TypeOf(arguments["FILE"]))

}

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

func prepareHostsList(urls []string) ([]string, error) {
	hosts := make([]string, 0)
	for i := range urls {
		h, err := getRemoteList(urls[i])
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		hosts = append(hosts, h...)
	}

	return hosts, nil
}
