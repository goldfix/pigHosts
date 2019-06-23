package pighosts

import (
	"reflect"
	"strings"
	"testing"
)

func init() {
	InitPigHosts(true)
	ReadFileConf()
}

func Test_removeComments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"remove: comment_1", args{s: "host.local.it #test"}, "host.local.it"},
		{"remove: comment_2", args{s: "host.local.it#test"}, "host.local.it"},
		{"remove: comment_3", args{s: "#host.local.it"}, ""},
		{"remove: comment_4", args{s: "# host.local.it #test"}, ""},
		{"remove: comment_space", args{s: "255.255.255.255                      broadcasthost"}, "255.255.255.255 broadcasthost"},
		{"remove: comment_tab", args{s: "255.255.255.255		broadcasthost"}, "255.255.255.255 broadcasthost"},
		{"remove: comment_tab_upper", args{s: "255.255.255.255		BROADCASTHOST"}, "255.255.255.255 broadcasthost"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeComments(tt.args.s); got != tt.want {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeLocalHost(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"remove: 127.0.0.1", args{s: "127.0.0.1 host.local.it"}, "host.local.it"},
		{"remove: 127.0.0.1 with spaces", args{s: "127.0.0.1 host.local.it                "}, "host.local.it"},
		{"remove: 0.0.0.0", args{s: "0.0.0.0 host.local.it"}, "host.local.it"},
		{"remove: 0.0.0.0 with spaces", args{s: "0.0.0.0     host.local.it"}, "host.local.it"},
		{"remove: 0.0.0.0 with spaces 2", args{s: "          0.0.0.0     host.local.it           "}, "host.local.it"},
		{"remove: 0.0.0.0 with localhost", args{s: "localhost"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLocalHost(tt.args.s); got != tt.want {
				t.Errorf("removeLocalHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareHostsList(t *testing.T) {

	tmpUrls := make([]string, 0)
	tmpUrl, _ := downlaodRemoteList("https://drive.google.com/uc?authuser=0&id=1-QRZf_ymrWFZ4XgmXTZJrkhqzhdJMphB&export=download")
	tmpUrls = append(tmpUrls, tmpUrl...)
	tmpUrl, _ = downlaodRemoteList("https://drive.google.com/uc?authuser=0&id=1BfGJJLtimhoOi9Sm3jYLF6d8XtYBJ5KY&export=download")
	tmpUrls = append(tmpUrls, tmpUrl...)

	type args struct {
		urls []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{"Test_prepareHostsList_1",
			args{urls: tmpUrls},
			map[string]int{
				"123date.me":             1,
				"12place.com":            1,
				"165a7.v.fwmrm.net":      1,
				"180searchassistant.com": 1,
				"188server.com":          1,
				"1ccbt.com":              1,
				"1empiredirect.com":      2,
				"1phads.com":             2,
				"test.test.io":           1},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prepareHostsList(tt.args.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareHostsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareHostsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitHostPerLine(t *testing.T) {
	type args struct {
		hosts map[string]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_splitHostPerLine_1",
			args{hosts: map[string]int{
				"123date.me":             1,
				"12place.com":            1,
				"165a7.v.fwmrm.net":      1,
				"180searchassistant.com": 1,
				"188server.com":          1,
				"1ccbt.com":              1,
				"1empiredirect.com":      2,
				"1phads.com":             2,
				"1phads2.com":            2,
				"1phads3.com":            2,
				"1phads4.com":            2,
				"1phads5.com":            2,
				"1phads6.com":            2,
				"1phads7.com":            2,
				"test.test.io":           1},
			},
			[]int{10, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitHostPerLine(tt.args.hosts)
			iGot := []int{len(strings.Split(got[0], " ")), len(strings.Split(got[1], " "))}

			if !reflect.DeepEqual(iGot, tt.want) {
				t.Errorf("splitHostPerLine() = %v, want %v", iGot, tt.want)
			}
		})
	}
}
